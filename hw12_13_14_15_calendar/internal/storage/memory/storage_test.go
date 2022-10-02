package memorystorage

import (
	"testing"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func generateDummyEvent(title string, desc string, addToDate time.Duration) (event storage.Event, err error) {
	UUID, err := uuid.NewUUID()
	if err != nil {
		return
	}

	userUUID, err := uuid.NewUUID()
	if err != nil {
		return
	}

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return
	}

	dummyDate := time.Date(2022, 01, 01, 0, 0, 0, 0, location)

	event = storage.Event{
		ID:                 UUID,
		Title:              title,
		DateStart:          dummyDate.Add(addToDate),
		Duration:           time.Hour * 24,
		Description:        desc,
		UserID:             userUUID,
		NotificationPeriod: time.Hour * 12,
	}

	return
}

func TestStorage(t *testing.T) {
	t.Run("Test add event", func(t *testing.T) {
		s := New()

		event, err := generateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		require.Equal(t, 1, len(s.events))
		require.Equal(t, event.ID, s.events[0].ID)
	})

	t.Run("Test update event", func(t *testing.T) {
		s := New()

		event, err := generateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		updatedEvent := storage.Event{
			ID:                 event.ID,
			Title:              "new title",
			DateStart:          event.DateStart,
			Duration:           event.Duration,
			Description:        event.Description,
			UserID:             event.UserID,
			NotificationPeriod: event.NotificationPeriod,
		}

		_, err = s.UpdateEventByID(updatedEvent.ID, updatedEvent)
		require.NoError(t, err)

		require.Equal(t, 1, len(s.events))
		require.Equal(t, "new title", s.events[0].Title)
	})

	t.Run("Test delete event", func(t *testing.T) {
		s := New()

		event, err := generateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		_, err = s.DeleteEvent(event.ID)
		require.NoError(t, err)

		require.Equal(t, 0, len(s.events))
	})

	t.Run("Test get by id event", func(t *testing.T) {
		s := New()

		event, err := generateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		eventByID, err := s.GetEventByID(event.ID)
		require.NoError(t, err)

		require.Equal(t, event.ID, eventByID.ID)
		require.Equal(t, event.Title, eventByID.Title)
	})

	t.Run("Test get events", func(t *testing.T) {
		s := New()

		event, err := generateDummyEvent("some title 1", "some description 1", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		event, err = generateDummyEvent("some title 2", "some description 2", time.Hour*24)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		events, err := s.GetEvents()
		require.NoError(t, err)

		require.Equal(t, 2, len(events))
	})

	t.Run("Test get events for specific period", func(t *testing.T) {
		s := New()

		event1, err := generateDummyEvent("some title 1", "some description 1", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event1)
		require.NoError(t, err)

		event2, err := generateDummyEvent("some title 2", "some description 2", time.Hour*24*2)
		require.NoError(t, err)

		_, err = s.AddEvent(event2)
		require.NoError(t, err)

		events, err := s.GetEventsForPeriod(event1.DateStart, time.Hour*24)
		require.NoError(t, err)

		require.Equal(t, 1, len(events))
	})
}
