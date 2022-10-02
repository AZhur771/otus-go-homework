package storage

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	// Уникальный идентификатор события (можно воспользоваться UUID)
	ID uuid.UUID `db:"id"`

	// Заголовок
	Title string `db:"title"`

	// Дата и время события
	DateStart time.Time `db:"date_start"`

	// Длительность события
	Duration time.Duration `db:"duration"`

	// Описание события
	Description string `db:"description"`

	// ID пользователя, владельца события
	UserID uuid.UUID `db:"user_id"`

	// За сколько времени высылать уведомление
	NotificationPeriod time.Duration `db:"notification_period"`
}
