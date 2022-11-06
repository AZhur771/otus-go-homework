package memorystorage

import (
	"testing"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Run("Test add event", func(t *testing.T) {
		s := New()

		event, err := storage.GenerateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		_, err = s.AddEvent(event)
		require.NoError(t, err)

		require.Equal(t, 1, len(s.events))
		require.Equal(t, event.Title, s.events[0].Title)
	})

	t.Run("Test update event", func(t *testing.T) {
		s := New()

		event, err := storage.GenerateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		s.events = append(s.events, event)

		updatedEvent := storage.Event{
			ID:                 event.ID,
			Title:              "new title",
			DateStart:          event.DateStart,
			Duration:           event.Duration,
			Description:        event.Description,
			UserID:             event.UserID,
			NotificationPeriod: event.NotificationPeriod,
		}

		_, err = s.UpdateEventByID(updatedEvent)
		require.NoError(t, err)

		require.Equal(t, 1, len(s.events))
		require.Equal(t, "new title", s.events[0].Title)
	})

	t.Run("Test delete event", func(t *testing.T) {
		s := New()

		event, err := storage.GenerateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		s.events = append(s.events, event)

		_, err = s.DeleteEventByID(event.ID)
		require.NoError(t, err)

		require.Equal(t, 0, len(s.events))
	})

	t.Run("Test get by id event", func(t *testing.T) {
		s := New()

		event, err := storage.GenerateDummyEvent("some title", "some description", 0)
		require.NoError(t, err)

		s.events = append(s.events, event)

		eventByID, err := s.GetEventByID(event.ID)
		require.NoError(t, err)

		require.Equal(t, event.ID, eventByID.ID)
		require.Equal(t, event.Title, eventByID.Title)
	})

	t.Run("Test get events", func(t *testing.T) {
		s := New()

		event1, err := storage.GenerateDummyEvent("some title 1", "some description 1", 0)
		require.NoError(t, err)

		event2, err := storage.GenerateDummyEvent("some title 2", "some description 2", storage.PqDuration(time.Hour*24))
		require.NoError(t, err)

		s.events = append(s.events, event1, event2)

		events, err := s.GetEvents()
		require.NoError(t, err)

		require.Equal(t, 2, len(events))
	})

	t.Run("Test get events for specific period", func(t *testing.T) {
		s := New()

		event1, err := storage.GenerateDummyEvent("some title 1", "some description 1", 0)
		require.NoError(t, err)

		event2, err := storage.GenerateDummyEvent("some title 2", "some description 2", storage.PqDuration(time.Hour*24*2))
		require.NoError(t, err)

		s.events = append(s.events, event1, event2)

		events, err := s.GetEventsForPeriod(event1.DateStart, event1.DateStart.Add(time.Hour*24))
		require.NoError(t, err)

		require.Equal(t, 1, len(events))
	})
}
