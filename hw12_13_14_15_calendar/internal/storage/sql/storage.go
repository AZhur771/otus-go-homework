package sqlstorage

import (
	"context"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db      *sqlx.DB
	timeout time.Duration
}

func New(timeout time.Duration) *Storage {
	return &Storage{
		timeout: timeout,
	}
}

func (s *Storage) Connect(ctx context.Context, datasource string, maxConnections int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	db, err := sqlx.ConnectContext(ctx, "postgres", datasource)
	db.SetMaxOpenConns(maxConnections)
	s.db = db
	return err
}

func (s *Storage) Close() error {
	err := s.db.Close()
	return err
}

func (s *Storage) AddEvent(event storage.Event) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	sql := `
		INSERT INTO events (title, date_start, duration, description, user_id, notification_period)
		VALUES (:title, :date_start, :duration, :description, :user_id, :notification_period)
		RETURNING *
	`

	row, err := s.db.NamedQueryContext(ctx, sql, event)
	if err != nil {
		return event, err
	}
	defer row.Close()

	row.Next()
	err = row.StructScan(&event)

	return event, err
}

func (s *Storage) DeleteEventByID(id uuid.UUID) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	var event storage.Event

	// Delete event and return this event data
	sql := "DELETE FROM events WHERE id = :id RETURNING *"

	row, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return event, err
	}
	defer row.Close()

	row.Next()
	err = row.StructScan(&event)

	return event, err
}

func (s *Storage) DeleteScheduledEvents(deletePeriod time.Time) ([]storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	deletedEvents := make([]storage.Event, 0)

	sql := "DELETE FROM events WHERE date_start <= :delete_period AND sent = true RETURNING *"

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"delete_period": deletePeriod,
	})
	if err != nil {
		return deletedEvents, err
	}
	defer rows.Close()

	for rows.Next() {
		var event storage.Event

		err := rows.StructScan(&event)
		if err != nil {
			return deletedEvents, err
		}

		deletedEvents = append(deletedEvents, event)
	}

	return deletedEvents, nil
}

func (s *Storage) UpdateEventByID(event storage.Event) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	sql := `
		UPDATE events
		SET title = :title,
		    date_start = :date_start,
		    duration = :duration,
		    description = :description,
		    user_id = :user_id,
		    notification_period = :notification_period
		WHERE id = :id
		RETURNING *
	`

	row, err := s.db.NamedQueryContext(ctx, sql, event)
	if err != nil {
		return event, err
	}
	defer row.Close()

	row.Next()
	err = row.StructScan(&event)

	return event, err
}

func (s *Storage) GetEventByID(id uuid.UUID) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	var event storage.Event

	sql := "SELECT * FROM events WHERE id = :id"

	row, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return event, err
	}
	defer row.Close()

	row.Next()
	err = row.StructScan(&event)

	return event, err
}

func (s *Storage) GetEvents() ([]storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	events := make([]storage.Event, 0)

	sql := "SELECT * FROM events"

	rows, err := s.db.QueryxContext(ctx, sql)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event storage.Event

		err := rows.StructScan(&event)
		if err != nil {
			return events, err
		}

		events = append(events, event)
	}

	return events, err
}

func (s *Storage) GetEventsForPeriod(startPeriod time.Time, endPeriod time.Time) ([]storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	events := make([]storage.Event, 0)

	sql := "SELECT * FROM events WHERE date_start >= :start AND date_start < :end"

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"start": startPeriod,
		"end":   endPeriod,
	})
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event storage.Event

		err := rows.StructScan(&event)
		if err != nil {
			return events, err
		}

		events = append(events, event)
	}

	return events, err
}

func (s *Storage) GetScheduledEvents(scanPeriod time.Time) ([]storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	events := make([]storage.Event, 0)

	sql := "SELECT * FROM events WHERE date_start - notification_period <= :scan_period AND sent = false"

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"scan_period": scanPeriod,
	})
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event storage.Event

		err := rows.StructScan(&event)
		if err != nil {
			return events, err
		}

		events = append(events, event)
	}

	return events, err
}

func (s *Storage) MarkEventsAsSent(ids []uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	query, args, err := sqlx.In("UPDATE events SET sent = true WHERE id IN (?)", ids)
	if err != nil {
		return err
	}

	query = s.db.Rebind(query)
	rows, err := s.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
