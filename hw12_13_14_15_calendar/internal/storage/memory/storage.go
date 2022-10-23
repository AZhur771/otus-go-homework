package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

var ErrNoEventFound = errors.New("event not found error")

type Storage struct {
	mu     sync.RWMutex
	events []storage.Event
}

func New() *Storage {
	return &Storage{
		mu:     sync.RWMutex{},
		events: make([]storage.Event, 0),
	}
}

func getEventAndEventIdx(id uuid.UUID, events []storage.Event) (storage.Event, int, error) {
	var event storage.Event
	eventIdx := -1

	for idx, e := range events {
		if e.ID == id {
			event = e
			eventIdx = idx
		}
	}

	if eventIdx == -1 {
		return event, eventIdx, ErrNoEventFound
	}

	return event, eventIdx, nil
}

func (s *Storage) AddEvent(event storage.Event) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return event, nil
}

func (s *Storage) DeleteEventByID(id uuid.UUID) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, eventIdx, err := getEventAndEventIdx(id, s.events)
	if err != nil {
		return event, err
	}

	s.events = append(s.events[:eventIdx], s.events[eventIdx+1:]...)

	return event, nil
}

func (s *Storage) UpdateEventByID(event storage.Event) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, eventIdx, err := getEventAndEventIdx(event.ID, s.events)
	if err != nil {
		return event, err
	}

	s.events[eventIdx] = event

	return event, nil
}

func (s *Storage) GetEventByID(id uuid.UUID) (storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	event, _, err := getEventAndEventIdx(id, s.events)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *Storage) GetEvents() ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.events, nil
}

func (s *Storage) GetEventsForPeriod(startPeriod time.Time, endPeriod time.Time) ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	eventsFiltered := make([]storage.Event, 0)

	for _, e := range s.events {
		if startPeriod.Equal(e.DateStart) || startPeriod.Before(e.DateStart) && endPeriod.After(e.DateStart) {
			eventsFiltered = append(eventsFiltered, e)
		}
	}

	return eventsFiltered, nil
}
