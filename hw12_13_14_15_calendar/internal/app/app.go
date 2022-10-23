package app

import (
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"time"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type Storage interface {
	AddEvent(event storage.Event) (storage.Event, error)
	DeleteEventById(id uuid.UUID) (storage.Event, error)
	UpdateEventByID(event storage.Event) (storage.Event, error)
	GetEventByID(id uuid.UUID) (storage.Event, error)
	GetEvents() ([]storage.Event, error)
	GetEventsForPeriod(dateStart time.Time, duration storage.PqDuration) ([]storage.Event, error)
}
