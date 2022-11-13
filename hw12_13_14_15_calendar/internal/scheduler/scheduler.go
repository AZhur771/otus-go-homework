package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
)

type Scheduler struct {
	storage        app.Storage
	logger         app.Logger
	producer       app.Producer
	scanPeriod     time.Duration
	deletePeriod   time.Duration
	startImmediate bool
}

func New(
	storage app.Storage,
	logger app.Logger,
	producer app.Producer,
	scanPeriod, deletePeriod time.Duration,
	startImmediate bool,
) *Scheduler {
	return &Scheduler{
		storage:        storage,
		logger:         logger,
		producer:       producer,
		scanPeriod:     scanPeriod,
		deletePeriod:   deletePeriod,
		startImmediate: startImmediate,
	}
}

func (s *Scheduler) ScanDatabaseForNotifier(ctx context.Context) (int, error) {
	// TODO: batch select events
	events, err := s.storage.GetScheduledEvents(time.Now().UTC())
	if err != nil {
		return 0, fmt.Errorf("failed to get events: %w", err)
	}

	sent := 0

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
			sent++
		}
	}

	return sent, nil
}

func (s *Scheduler) processScanDatabaseForNotifierResult(sent int, err error) {
	if err != nil {
		s.logger.Error(fmt.Sprintf("error while scanning database: %s", err))
	} else {
		s.logger.Info(fmt.Sprintf("database scanned, events sent: %d", sent))
	}
}

func (s *Scheduler) RunNotifier(ctx context.Context) error {
	t := time.NewTicker(s.scanPeriod)
	defer t.Stop()
	for {
		if s.startImmediate {
			sent, err := s.ScanDatabaseForNotifier(ctx)
			s.processScanDatabaseForNotifierResult(sent, err)
		}

		select {
		case <-t.C:
			sent, err := s.ScanDatabaseForNotifier(ctx)
			s.processScanDatabaseForNotifierResult(sent, err)
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
