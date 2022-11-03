//nolint:typecheck,nolintlint
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
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func generateDummyEvent(title string, desc string, addToDate storage.PqDuration) (event storage.Event, err error) {
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

	event = storage.Event{
		ID:                 UUID,
		Title:              title,
		DateStart:          dummyDate.Add(time.Duration(addToDate)),
		Duration:           storage.PqDuration(time.Hour * 24),
		Description:        desc,
		UserID:             userUUID,
		NotificationPeriod: storage.PqDuration(time.Hour * 12),
	}

	return
}

type ServerTestSuite struct {
	storage app.Storage
	server  eventpb.EventServiceServer
	suite.Suite
}

func (s *ServerTestSuite) SetupTest() {
	s.storage = memoryStorage.New()
	s.server = NewServer(s.storage, logger.New(logger.Debug))
}

func (s *ServerTestSuite) TestAddEvent() {
	UUID, err := uuid.NewUUID()
	s.Require().NoError(err)

	req := &eventpb.AddEventRequest{
		UserId:             UUID.String(),
		Title:              "some title",
		Description:        "some description",
		DateStart:          timestamppb.New(time.Now()),
		Duration:           durationpb.New(time.Hour),
		NotificationPeriod: durationpb.New(time.Hour),
	}

	res, err := s.server.AddEvent(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	events, err := s.storage.GetEvents()
	s.Require().NoError(err)
	s.Require().Equal(1, len(events))
}

func (s *ServerTestSuite) TestUpdateEvent() {
	event, err := generateDummyEvent("some title", "some description", 0)
	s.Require().NoError(err)
	s.storage.AddEvent(event)

	events, err := s.storage.GetEvents()
	s.Require().NoError(err)

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

	res, err := s.server.UpdateEventByID(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	event, err = s.storage.GetEventByID(eventID)
	s.Require().NoError(err)
	s.Require().Equal(req.Title, event.Title)
	s.Require().Equal(req.Description, event.Description)
}

func (s *ServerTestSuite) TestGetEvent() {
	event, err := generateDummyEvent("some title", "some description", 0)
	s.Require().NoError(err)
	s.storage.AddEvent(event)

	events, err := s.storage.GetEvents()
	s.Require().NoError(err)

	eventID := events[0].ID

	req := &eventpb.GetEventByIDRequest{
		Id: eventID.String(),
	}

	res, err := s.server.GetEventByID(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(res.Title, event.Title)
}

func (s *ServerTestSuite) TestDeleteEvent() {
	event, err := generateDummyEvent("some title", "some description", 0)
	s.Require().NoError(err)
	s.storage.AddEvent(event)

	events, err := s.storage.GetEvents()
	s.Require().NoError(err)

	eventID := events[0].ID

	req := &eventpb.DeleteEventByIDRequest{
		Id: eventID.String(),
	}

	res, err := s.server.DeleteEventByID(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	events, err = s.storage.GetEvents()
	s.Require().NoError(err)
	s.Require().Equal(0, len(events))
}

func (s *ServerTestSuite) TestGetEvents() {
	for i := 0; i < 3; i++ {
		event, err := generateDummyEvent(fmt.Sprintf("some title %d", i), fmt.Sprintf("some description %d", i), 0)
		s.Require().NoError(err)

		s.storage.AddEvent(event)
	}

	req := &eventpb.GetEventsRequest{}

	res, err := s.server.GetEvents(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(3, len(res.Events))
}

func (s *ServerTestSuite) TestGetEventsForPeriod() {
	event1, err := generateDummyEvent("some title 1", "some description 1", 0)
	s.Require().NoError(err)

	s.storage.AddEvent(event1)

	event2, err := generateDummyEvent("some title 2", "some description 2", storage.PqDuration(time.Hour*24*2))
	s.Require().NoError(err)

	s.storage.AddEvent(event2)

	req := &eventpb.GetEventsRequest{
		PeriodStart: timestamppb.New(event1.DateStart),
		PeriodEnd:   timestamppb.New(event1.DateStart.Add(time.Hour)),
	}

	res, err := s.server.GetEvents(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(1, len(res.Events))
	s.Require().Equal("some title 1", res.Events[0].Title)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
