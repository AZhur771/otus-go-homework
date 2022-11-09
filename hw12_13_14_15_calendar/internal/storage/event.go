package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	// Уникальный идентификатор события (можно воспользоваться UUID)
	ID uuid.UUID `db:"id"`

	// Заголовок
	Title string `db:"title"`

	// Дата и время события
	DateStart time.Time `db:"date_start"`

	// Длительность события
	Duration PqDuration `db:"duration"`

	// Описание события
	Description string `db:"description"`

	// ID пользователя, владельца события
	UserID uuid.UUID `db:"user_id"`

	// За сколько времени высылать уведомление
	NotificationPeriod PqDuration `db:"notification_period"`

	// Было ли сообщение отправлено
	Sent bool `db:"sent"`
}
