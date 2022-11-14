package sender

import (
	"context"
	"fmt"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
)

type Sender struct {
	storage  app.Storage
	logger   app.Logger
	consumer app.Consumer
}

func New(storage app.Storage, logger app.Logger, consumer app.Consumer) *Sender {
	return &Sender{
		storage:  storage,
		logger:   logger,
		consumer: consumer,
	}
}

func (s *Sender) sendNotification(msg app.Message) error {
	// There should be some sending notification logic
	s.logger.Info(fmt.Sprintf("Received message %s for user %s: %s - %s",
		msg.ID, msg.UserID, msg.Title, msg.Date))

	return nil
}

func (s *Sender) Run(ctx context.Context) error {
	msgs, err := s.consumer.Consume(ctx)
	if err != nil {
		return fmt.Errorf("error while consuming messages: %w", err)
	}

	for msg := range msgs {
		if err = s.sendNotification(msg); err != nil {
			continue
		}

		if err = s.storage.MarkEventAsSent(msg.ID); err != nil {
			s.logger.Error(fmt.Sprintf("failed to mark events as sent: %s", err))
		}
	}

	return nil
}
