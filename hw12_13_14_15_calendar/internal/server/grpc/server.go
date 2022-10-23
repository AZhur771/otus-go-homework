package internalgrpc

import (
	"context"
	"errors"
	"time"

	eventpb "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/api/stubs"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrPeriodStartAbsent = errors.New("periodStart is not passed")
	ErrPeriodEndAbsent   = errors.New("periodEnd is not passed")
)

type EventServiceServerImpl struct {
	eventpb.UnimplementedEventServiceServer
	storage app.Storage
	logger  app.Logger
}

func NewServer(storage app.Storage, logger app.Logger) eventpb.EventServiceServer {
	return EventServiceServerImpl{
		storage: storage,
		logger:  logger,
	}
}

func (e EventServiceServerImpl) AddEvent(
	ctx context.Context,
	request *eventpb.AddEventRequest,
) (*eventpb.Event, error) {
	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, err
	}

	event := storage.Event{
		UserID:             userID,
		Title:              request.GetTitle(),
		Description:        request.GetDescription(),
		DateStart:          request.GetDateStart().AsTime(),
		Duration:           storage.PqDuration(request.GetDuration().AsDuration()),
		NotificationPeriod: storage.PqDuration(request.GetNotificationPeriod().AsDuration()),
	}

	event, err = e.storage.AddEvent(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.Event{
		Id:                 event.ID.String(),
		UserId:             event.UserID.String(),
		Title:              event.Title,
		Description:        event.Description,
		DateStart:          timestamppb.New(event.DateStart),
		Duration:           durationpb.New(time.Duration(event.Duration)),
		NotificationPeriod: durationpb.New(time.Duration(event.NotificationPeriod)),
	}, nil
}

func (e EventServiceServerImpl) DeleteEventByID(
	ctx context.Context,
	request *eventpb.DeleteEventByIDRequest,
) (*eventpb.Event, error) {
	eventID, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	event, err := e.storage.DeleteEventByID(eventID)
	if err != nil {
		return nil, err
	}

	return &eventpb.Event{
		Id:                 event.ID.String(),
		UserId:             event.UserID.String(),
		Title:              event.Title,
		Description:        event.Description,
		DateStart:          timestamppb.New(event.DateStart),
		Duration:           durationpb.New(time.Duration(event.Duration)),
		NotificationPeriod: durationpb.New(time.Duration(event.NotificationPeriod)),
	}, nil
}

func (e EventServiceServerImpl) UpdateEventByID(
	ctx context.Context,
	request *eventpb.Event,
) (*eventpb.Event, error) {
	eventID, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, err
	}

	event := storage.Event{
		ID:                 eventID,
		UserID:             userID,
		Title:              request.GetTitle(),
		Description:        request.GetDescription(),
		DateStart:          request.GetDateStart().AsTime(),
		Duration:           storage.PqDuration(request.GetDuration().AsDuration()),
		NotificationPeriod: storage.PqDuration(request.GetNotificationPeriod().AsDuration()),
	}

	event, err = e.storage.UpdateEventByID(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.Event{
		Id:                 event.ID.String(),
		UserId:             event.UserID.String(),
		Title:              event.Title,
		Description:        event.Description,
		DateStart:          timestamppb.New(event.DateStart),
		Duration:           durationpb.New(time.Duration(event.Duration)),
		NotificationPeriod: durationpb.New(time.Duration(event.NotificationPeriod)),
	}, nil
}

func (e EventServiceServerImpl) GetEventByID(
	ctx context.Context,
	request *eventpb.GetEventByIDRequest,
) (*eventpb.Event, error) {
	eventID, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	event, err := e.storage.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	return &eventpb.Event{
		Id:                 event.ID.String(),
		UserId:             event.UserID.String(),
		Title:              event.Title,
		Description:        event.Description,
		DateStart:          timestamppb.New(event.DateStart),
		Duration:           durationpb.New(time.Duration(event.Duration)),
		NotificationPeriod: durationpb.New(time.Duration(event.NotificationPeriod)),
	}, nil
}

func (e EventServiceServerImpl) GetEvents(
	ctx context.Context,
	request *eventpb.GetEventsRequest,
) (*eventpb.Events, error) {
	periodStart := request.GetPeriodStart()
	periodEnd := request.GetPeriodEnd()

	var events []storage.Event
	var err error

	if periodStart != nil || periodEnd != nil {
		if periodStart == nil {
			return nil, ErrPeriodStartAbsent
		}

		if periodEnd == nil {
			return nil, ErrPeriodEndAbsent
		}

		events, err = e.storage.GetEventsForPeriod(periodStart.AsTime(), periodEnd.AsTime())
	} else {
		events, err = e.storage.GetEvents()
	}
	if err != nil {
		return nil, err
	}

	eventspb := make([]*eventpb.Event, 0)

	for _, event := range events {
		eventspb = append(eventspb, &eventpb.Event{
			Id:                 event.ID.String(),
			UserId:             event.UserID.String(),
			Title:              event.Title,
			Description:        event.Description,
			DateStart:          timestamppb.New(event.DateStart),
			Duration:           durationpb.New(time.Duration(event.Duration)),
			NotificationPeriod: durationpb.New(time.Duration(event.NotificationPeriod)),
		})
	}

	return &eventpb.Events{Events: eventspb}, nil
}
