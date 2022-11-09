package app

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
	UserID uuid.UUID `json:"userId"`
}
