package app

import (
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

type Storage interface {
	AddEvent(event storage.Event) (storage.Event, error)
	DeleteEventByID(id uuid.UUID) (storage.Event, error)
	UpdateEventByID(event storage.Event) (storage.Event, error)
	GetEventByID(id uuid.UUID) (storage.Event, error)
	GetEvents() ([]storage.Event, error)
	GetEventsForPeriod(periodStart time.Time, periodEnd time.Time) ([]storage.Event, error)
}
