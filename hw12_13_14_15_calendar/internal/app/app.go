package app

import (
	"context"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type App struct {
	storage Storage
	logger  Logger
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type Storage interface {
	AddEvent(event storage.Event) (storage.Event, error)
	DeleteEvent(id uuid.UUID) (storage.Event, error)
	UpdateEventByID(event storage.Event) (storage.Event, error)
	GetEventByID(id uuid.UUID) (storage.Event, error)
	GetEvents() ([]storage.Event, error)
	GetEventsForPeriod(dateStart time.Time, duration time.Duration) ([]storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{
		storage: storage,
		logger:  logger,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
