package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
)

type Scheduler struct {
	storage      app.Storage
	logger       app.Logger
	producer     app.Producer
	scanPeriod   time.Duration
	deletePeriod time.Duration
}

func New(
	storage app.Storage,
	logger app.Logger,
	producer app.Producer,
	scanPeriod, deletePeriod time.Duration,
) *Scheduler {
	return &Scheduler{
		storage:      storage,
		logger:       logger,
		producer:     producer,
		scanPeriod:   scanPeriod,
		deletePeriod: deletePeriod,
	}
}

func (s *Scheduler) ScanDatabase(ctx context.Context) (int, int, error) {
	events, err := s.storage.GetEvents()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get events: %w", err)
	}

	sent := 0
	deleted := 0

	startDate := time.Now().UTC()
	endDate := startDate.Add(s.scanPeriod)

	deleteDate := startDate.Add(-s.deletePeriod)

	for _, e := range events {
		notifyDate := e.DateStart.Add(-time.Duration(e.NotificationPeriod)).UTC()

		// Is event notification to be sent?
		if (notifyDate.After(startDate) || notifyDate.Equal(startDate)) && notifyDate.Before(endDate) {
			msg := app.Message{
				ID:     e.ID,
				Title:  e.Title,
				Date:   e.DateStart,
				UserID: e.UserID,
			}
			err := s.producer.Publish(ctx, msg)
			if err != nil {
				s.logger.Error(fmt.Sprintf("Failed to publish event with id %s: %s", e.ID.String(), err))
			} else {
				sent++
			}
		}

		// Is event to be deleted?
		if e.DateStart.Before(deleteDate) || e.DateStart.Equal(deleteDate) {
			_, err := s.storage.DeleteEventByID(e.ID)
			if err != nil {
				s.logger.Error(fmt.Sprintf("Failed to delete event with id %s: %s", e.ID.String(), err))
			} else {
				deleted++
			}
		}
	}

	return sent, deleted, nil
}

func (s *Scheduler) Run(ctx context.Context) error {
	t := time.NewTicker(s.scanPeriod)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			sent, deleted, err := s.ScanDatabase(ctx)
			if err != nil {
				s.logger.Error(fmt.Sprintf("error while scanning database: %s", err))
			} else {
				s.logger.Info(fmt.Sprintf("database scanned: %d notifications sent, %d events deleted", sent, deleted))
			}
		case <-ctx.Done():
			s.logger.Info("exit gracefully")
			return nil
		}
	}
}
