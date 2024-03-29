package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	uuid "github.com/google/uuid"
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

func contains(ids []uuid.UUID, id uuid.UUID) bool {
	for _, v := range ids {
		if v.String() == id.String() {
			return true
		}
	}
	return false
}

func (s *Storage) AddEvent(event storage.Event) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	UUID, err := uuid.NewUUID()
	if err != nil {
		return event, err
	}
	event.ID = UUID
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

func (s *Storage) DeleteScheduledEvents(deletePeriod time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	deletedEvents := make([]storage.Event, 0)
	filteredEvents := make([]storage.Event, 0)

	for _, e := range s.events {
		if (deletePeriod.Equal(e.DateStart) || deletePeriod.After(e.DateStart)) && e.Sent {
			deletedEvents = append(deletedEvents, e)
		} else {
			filteredEvents = append(filteredEvents, e)
		}
	}

	s.events = filteredEvents

	return deletedEvents, nil
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
		if (startPeriod.Equal(e.DateStart) || startPeriod.Before(e.DateStart)) && endPeriod.After(e.DateStart) {
			eventsFiltered = append(eventsFiltered, e)
		}
	}

	return eventsFiltered, nil
}

func (s *Storage) GetScheduledEvents(scanPeriod time.Time) ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	scheduledEvents := make([]storage.Event, 0)

	for _, e := range s.events {
		notificationDate := e.DateStart.Add(-time.Duration(e.NotificationPeriod))
		if (scanPeriod.Equal(notificationDate) || scanPeriod.After(notificationDate)) && !e.Sent {
			scheduledEvents = append(scheduledEvents, e)
		}
	}

	return scheduledEvents, nil
}

func (s *Storage) MarkEventAsSent(id uuid.UUID) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, e := range s.events {
		if id.String() == e.ID.String() {
			s.events[i].Sent = true
			break
		}
	}

	return nil
}

func (s *Storage) MarkEventsAsSent(ids []uuid.UUID) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, e := range s.events {
		if contains(ids, e.ID) {
			s.events[i].Sent = true
		}
	}

	return nil
}
