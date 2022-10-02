package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct { // TODO
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, host string, port int, username, password, dbname, sslmode string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, dbname, sslmode)
	db, err := sqlx.ConnectContext(ctx, "postgres", datasource)
	db.SetMaxOpenConns(5)
	s.db = db
	return err
}

func (s *Storage) Close(ctx context.Context) error {
	err := s.db.Close()
	return err
}

func (s *Storage) AddEvent(event storage.Event) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := "INSERT INTO events (title, date_start, duration, description, user_id, notification_period) VALUES (:title, :date_start, :duration, :description, :user_id, :notification_period)"

	row, err := s.db.NamedQueryContext(ctx, sql, event)
	if err != nil {
		return event, err
	}

	err = row.Scan(event)
	return event, err
}

func (s *Storage) DeleteEvent(id uuid.UUID) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event storage.Event

	sql := "DELETE FROM events WHERE id = :id RETURNING *"

	row, err := s.db.NamedQueryContext(ctx, sql, event)
	if err != nil {
		return event, err
	}

	err = row.Scan(event)

	return event, err
}

func (s *Storage) UpdateEventByID(id uuid.UUID, event storage.Event) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := "UPDATE events SET title = :title, date_start = :date_start, duration = :duration, description = :description, user_id = :user_id, notification_period = :notification_period WHERE id = :id"

	row, err := s.db.NamedQueryContext(ctx, sql, event)
	if err != nil {
		return event, err
	}

	err = row.Scan(event)

	return event, err
}

func (s *Storage) GetEventByID(id uuid.UUID) (storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event storage.Event

	sql := "SELECT * FROM events WHERE id = :id"

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]uuid.UUID{
		"id": id,
	})
	if err != nil {
		return event, err
	}

	err = rows.Scan(event)
	return event, err
}

func (s *Storage) GetEvents() ([]storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	events := make([]storage.Event, 0)

	sql := "SELECT * FROM events"

	rows, err := s.db.QueryxContext(ctx, sql)
	if err != nil {
		return events, err
	}

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

func (s *Storage) GetEventsForPeriod(dateStart time.Time, duration time.Duration) ([]storage.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	events := make([]storage.Event, 0)

	sql := "SELECT * FROM events WHERE date_start >= :start AND date_start + duration < :end"

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]int64{
		"start": dateStart.Unix(),
		"end":   dateStart.Add(duration).Unix(),
	})
	if err != nil {
		return events, err
	}

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
