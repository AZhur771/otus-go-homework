package storage

import (
	"time"

	"github.com/google/uuid"
)

func GenerateDummyEvent(title string, desc string, addToDate PqDuration) (event Event, err error) {
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

	dummyDate := time.Date(2022, 1, 1, 0, 0, 0, 0, location)

	event = Event{
		ID:                 UUID,
		Title:              title,
		DateStart:          dummyDate.Add(time.Duration(addToDate)),
		Duration:           PqDuration(time.Hour * 24),
		Description:        desc,
		UserID:             userUUID,
		NotificationPeriod: PqDuration(time.Hour * 12),
	}

	return
}
