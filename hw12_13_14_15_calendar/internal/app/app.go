package app

import (
	"context"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type Producer interface {
	Connect(ctx context.Context, exchangeName, exchangeType,
		queueName, bindingKey string, persistent, reliable bool) error
	Disconnect() error
	Publish(ctx context.Context, msg Message) error
	WaitForConfirms(ctx context.Context) error
}

type Consumer interface {
	Connect(ctx context.Context, exchangeName, exchangeType,
		queueName, consumerTag, bindingKey string, persistent bool) error
	Disconnect() error
	Consume(ctx context.Context) (<-chan Message, error)
}

type Storage interface {
	AddEvent(event storage.Event) (storage.Event, error)
	DeleteEventByID(id uuid.UUID) (storage.Event, error)
	DeleteScheduledEvents(deletePeriod time.Time) ([]storage.Event, error)
	UpdateEventByID(event storage.Event) (storage.Event, error)
	GetEventByID(id uuid.UUID) (storage.Event, error)
	GetEvents() ([]storage.Event, error)
	GetEventsForPeriod(periodStart time.Time, periodEnd time.Time) ([]storage.Event, error)
	GetScheduledEvents(scanPeriod time.Time) ([]storage.Event, error)
	MarkEventAsSent(id uuid.UUID) error
	MarkEventsAsSent(ids []uuid.UUID) error
}
