package memorystorage

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
)

var NoEventFoundErr = errors.New("event not found error")

type Storage struct {
	mu     sync.RWMutex
	events []storage.Event
}

func New() *Storage {
	return &Storage{
		mu:     *new(sync.RWMutex),
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
		return event, eventIdx, NoEventFoundErr
	}

	return event, eventIdx, nil
}

func (s *Storage) AddEvent(event storage.Event) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return event, nil
}

func (s *Storage) DeleteEvent(id uuid.UUID) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, eventIdx, err := getEventAndEventIdx(id, s.events)
	if err != nil {
		return event, err
	}

	s.events = append(s.events[:eventIdx], s.events[eventIdx+1:]...)

	return event, nil
}

func (s *Storage) UpdateEventByID(id uuid.UUID, event storage.Event) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, eventIdx, err := getEventAndEventIdx(id, s.events)
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

func (s *Storage) GetEventsForPeriod(dateStart time.Time, duration time.Duration) ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	eventsFiltered := make([]storage.Event, 0)

	dateEnd := dateStart.Add(duration)

	for _, e := range s.events {
		if dateStart.Equal(e.DateStart) || dateStart.Before(e.DateStart) && dateEnd.After(e.DateStart) {
			eventsFiltered = append(eventsFiltered, e)
		}
	}

	return eventsFiltered, nil
}
