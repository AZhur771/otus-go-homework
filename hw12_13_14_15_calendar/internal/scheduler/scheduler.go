package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/google/uuid"
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

func (s *Scheduler) ScanDatabaseForNotifier(ctx context.Context) (int, error) {
	events, err := s.storage.GetScheduledEvents(time.Now().UTC())
	if err != nil {
		return 0, fmt.Errorf("failed to get events: %w", err)
	}

	sentEventIds := make([]uuid.UUID, 0)

	for _, e := range events {
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
			sentEventIds = append(sentEventIds, e.ID)
		}
	}

	err = s.storage.MarkEventsAsSent(sentEventIds)
	if err != nil {
		return 0, fmt.Errorf("failed to mark events as sent: %w", err)
	}

	return len(sentEventIds), nil
}

func (s *Scheduler) RunNotifier(ctx context.Context) error {
	t := time.NewTicker(s.scanPeriod)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			sent, err := s.ScanDatabaseForNotifier(ctx)
			if err != nil {
				s.logger.Error(fmt.Sprintf("error while scanning database: %s", err))
			} else {
				s.logger.Info(fmt.Sprintf("database scanned, events sent: %d", sent))
			}
		case <-ctx.Done():
			s.logger.Info("exit notifier gracefully")
			return nil
		}
	}
}

func (s *Scheduler) RunDeleter(ctx context.Context) error {
	t := time.NewTicker(s.scanPeriod)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			deleteDate := time.Now().UTC().Add(-s.deletePeriod)
			deletedEvents, err := s.storage.DeleteScheduledEvents(deleteDate)
			if err != nil {
				s.logger.Error(fmt.Sprintf("error while deleting events: %s", err))
			}
			s.logger.Info(fmt.Sprintf("database scanned, events deleted: %d", len(deletedEvents)))
		case <-ctx.Done():
			s.logger.Info("exit deleter gracefully")
			return nil
		}
	}
}
