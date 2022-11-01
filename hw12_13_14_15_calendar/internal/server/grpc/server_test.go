package internalgrpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	eventpb "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/api/stubs"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	memoryStorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerTestSuite struct {
	storage app.Storage
	server  eventpb.EventServiceServer
	suite.Suite
}

func (suit *ServerTestSuite) SetupTest() {
	suit.storage = memoryStorage.New()
	suit.server = NewServer(suit.storage, logger.New(logger.Debug))
}

func (suit *ServerTestSuite) TestAddEvent() {
	UUID, err := uuid.NewUUID()
	require.NoError(suit.T(), err)

	req := &eventpb.AddEventRequest{
		UserId:             UUID.String(),
		Title:              "some title",
		Description:        "some description",
		DateStart:          timestamppb.New(time.Now()),
		Duration:           durationpb.New(time.Hour),
		NotificationPeriod: durationpb.New(time.Hour),
	}

	res, err := suit.server.AddEvent(context.Background(), req)
	require.NoError(suit.T(), err)
	require.NotNil(suit.T(), res)

	events, err := suit.storage.GetEvents()
	require.NoError(suit.T(), err)
	require.Equal(suit.T(), 1, len(events))
}

func (suit *ServerTestSuite) TestUpdateEvent() {
	event, err := storage.GenerateDummyEvent("some title", "some description", 0)
	require.NoError(suit.T(), err)
	suit.storage.AddEvent(event)

	events, err := suit.storage.GetEvents()
	require.NoError(suit.T(), err)

	eventID := events[0].ID
	userID := events[0].UserID

	req := &eventpb.Event{
		Id:                 eventID.String(),
		UserId:             userID.String(),
		Title:              "new title",
		Description:        "new description",
		DateStart:          timestamppb.New(time.Now()),
		Duration:           durationpb.New(time.Hour),
		NotificationPeriod: durationpb.New(time.Hour),
	}

	res, err := suit.server.UpdateEventByID(context.Background(), req)
	require.NoError(suit.T(), err)
	require.NotNil(suit.T(), res)

	event, err = suit.storage.GetEventByID(eventID)
	require.NoError(suit.T(), err)
	require.Equal(suit.T(), req.Title, event.Title)
	require.Equal(suit.T(), req.Description, event.Description)
}

func (suit *ServerTestSuite) TestGetEvent() {
	event, err := storage.GenerateDummyEvent("some title", "some description", 0)
	require.NoError(suit.T(), err)
	suit.storage.AddEvent(event)

	events, err := suit.storage.GetEvents()
	require.NoError(suit.T(), err)

	eventID := events[0].ID

	req := &eventpb.GetEventByIDRequest{
		Id: eventID.String(),
	}

	res, err := suit.server.GetEventByID(context.Background(), req)
	require.NoError(suit.T(), err)
	require.NotNil(suit.T(), res)
	require.Equal(suit.T(), res.Title, event.Title)
}

func (suit *ServerTestSuite) TestDeleteEvent() {
	event, err := storage.GenerateDummyEvent("some title", "some description", 0)
	require.NoError(suit.T(), err)
	suit.storage.AddEvent(event)

	events, err := suit.storage.GetEvents()
	require.NoError(suit.T(), err)

	eventID := events[0].ID

	req := &eventpb.DeleteEventByIDRequest{
		Id: eventID.String(),
	}

	res, err := suit.server.DeleteEventByID(context.Background(), req)
	require.NoError(suit.T(), err)
	require.NotNil(suit.T(), res)

	events, err = suit.storage.GetEvents()
	require.NoError(suit.T(), err)
	require.Equal(suit.T(), 0, len(events))
}

func (suit *ServerTestSuite) TestGetEvents() {
	for i := 0; i < 3; i++ {
		event, err := storage.GenerateDummyEvent(fmt.Sprintf("some title %d", i), fmt.Sprintf("some description %d", i), 0)
		require.NoError(suit.T(), err)

		suit.storage.AddEvent(event)
	}

	req := &eventpb.GetEventsRequest{}

	res, err := suit.server.GetEvents(context.Background(), req)
	require.NoError(suit.T(), err)
	require.NotNil(suit.T(), res)
	require.Equal(suit.T(), 3, len(res.Events))
}

func (suit *ServerTestSuite) TestGetEventsForPeriod() {
	event1, err := storage.GenerateDummyEvent("some title 1", "some description 1", 0)
	require.NoError(suit.T(), err)

	suit.storage.AddEvent(event1)

	event2, err := storage.GenerateDummyEvent("some title 2", "some description 2", storage.PqDuration(time.Hour*24*2))
	require.NoError(suit.T(), err)

	suit.storage.AddEvent(event2)

	req := &eventpb.GetEventsRequest{
		PeriodStart: timestamppb.New(event1.DateStart),
		PeriodEnd:   timestamppb.New(event1.DateStart.Add(time.Hour)),
	}

	res, err := suit.server.GetEvents(context.Background(), req)
	require.NoError(suit.T(), err)
	require.NotNil(suit.T(), res)
	require.Equal(suit.T(), 1, len(res.Events))
	require.Equal(suit.T(), "some title 1", res.Events[0].Title)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
