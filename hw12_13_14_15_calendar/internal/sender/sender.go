package sender

import (
	"context"
	"fmt"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
)

type Sender struct {
	logger   app.Logger
	consumer app.Consumer
}

func New(logger app.Logger, consumer app.Consumer) *Sender {
	return &Sender{
		logger:   logger,
		consumer: consumer,
	}
}

func (s *Sender) Run(ctx context.Context) error {
	msgs, err := s.consumer.Consume(ctx)
	if err != nil {
		return fmt.Errorf("error while consuming messages: %w", err)
	}

	for msg := range msgs {
		s.logger.Info(fmt.Sprintf("Received message %s for user %s: %s - %s",
			msg.ID, msg.UserID, msg.Title, msg.Date))
	}

	return nil
}
