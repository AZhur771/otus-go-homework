// +build integration

package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage"
	"github.com/cucumber/messages-go/v16"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	eventpb "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/api/stubs"
	internalgrpc "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/server/grpc"
	"github.com/cucumber/godog"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const preparedUserId = "19a12b49-a57a-4f1e-8e66-152be08e6165"

var pqDatasource = os.Getenv("TEST_DATASOURCE")

var hostAndPortRegex = regexp.MustCompile(`grpc:\/\/(.+):(\d+)`)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type eventsTest struct {
	conn *grpc.ClientConn
	db   *sqlx.DB

	// use these fields to persist data between steps
	addedEventIds      []uuid.UUID
	deletedEventsIds   []uuid.UUID
	sentData           []byte
	eventsReturned     []storage.Event
	responseStatusCode codes.Code
}

func (test *eventsTest) setupTests(ctx context.Context, _ *messages.Pickle) (context.Context, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", pqDatasource)
	panicOnErr(err)
	db.SetMaxOpenConns(5)

	test.db = db

	// Prepopulate database with relevant events
	sql := `
		INSERT INTO events (title, date_start, duration, description, user_id, notification_period)
		VALUES (:title, :date_start, :duration, :description, :user_id, :notification_period)
		RETURNING *
	`

	userId, err := uuid.Parse(preparedUserId)
	panicOnErr(err)

	testEvents := []storage.Event{
		{
			Title:              "some title 1",
			DateStart:          time.Now().Add(10 * time.Minute).UTC(),
			Duration:           storage.PqDuration(2 * time.Hour),
			Description:        "some description 1",
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(0 * time.Hour),
		},
		{
			Title:              "some title 2",
			DateStart:          time.Now().Add(1 * time.Hour).UTC(),
			Duration:           storage.PqDuration(1 * time.Hour),
			Description:        "some description 2",
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(0 * time.Hour),
		},
		{
			Title:              "some title 3",
			DateStart:          time.Now().Add(2 * time.Hour).UTC(),
			Duration:           storage.PqDuration(2 * time.Hour),
			Description:        "some description 3",
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(0 * time.Hour),
		},
		{
			Title:              "some title 4",
			DateStart:          time.Now().Add(24 * time.Hour).UTC(),
			Duration:           storage.PqDuration(2 * time.Hour),
			Description:        "some description 4",
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(0 * time.Hour),
		},
	}

	for _, testEvent := range testEvents {
		row, err := db.NamedQueryContext(ctx, sql, testEvent)
		panicOnErr(err)

		var createdEvent storage.Event

		row.Next()
		err = row.StructScan(&createdEvent)
		panicOnErr(err)
		row.Close()

		test.addedEventIds = append(test.addedEventIds, createdEvent.ID)
	}

	return ctx, nil
}

func (test *eventsTest) cleanupTests(ctx context.Context, _ *messages.Pickle, _ error) (context.Context, error) {
	test.addedEventIds = make([]uuid.UUID, 0)
	test.deletedEventsIds = make([]uuid.UUID, 0)
	test.sentData = make([]byte, 0)
	test.eventsReturned = make([]storage.Event, 0)
	test.responseStatusCode = 0

	_, err := test.db.ExecContext(ctx, "DELETE FROM events")
	return ctx, err
}

func (test *eventsTest) setupTestConn(addr string) {
	hostAndPort := hostAndPortRegex.FindStringSubmatch(addr)
	host := hostAndPort[1]
	port, err := strconv.Atoi(hostAndPort[2])
	panicOnErr(err)

	test.conn, err = internalgrpc.NewClientConn(
		context.Background(),
		host,
		port,
	)
	panicOnErr(err)
}

func (test *eventsTest) iSendAddEventRequestToWithData(addr string, msg *godog.DocString) error {
	test.setupTestConn(addr)

	client := eventpb.NewEventServiceClient(test.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	protoMsg := &eventpb.AddEventRequest{}
	err := jsonpb.Unmarshal(strings.NewReader(msg.Content), protoMsg)
	panicOnErr(err)

	resp, err := client.AddEvent(ctx, protoMsg)
	if err != nil {
		return fmt.Errorf("failed to add event: %w", err)
	}

	UUID, err := uuid.Parse(resp.GetId())
	if err != nil {
		return fmt.Errorf("returned id is not a valid uuid: %w", err)
	}

	test.addedEventIds = append(test.addedEventIds, UUID)
	test.responseStatusCode = status.Code(err)

	return nil
}

func (test *eventsTest) iSendDeleteEventByIdRequestToWithIdOfTheLastCreatedEvent(addr string) error {
	test.setupTestConn(addr)

	client := eventpb.NewEventServiceClient(test.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	lastAddedEventId := test.addedEventIds[len(test.addedEventIds)-1]

	protoMsg := &eventpb.DeleteEventByIDRequest{Id: lastAddedEventId.String()}

	_, err := client.DeleteEventByID(ctx, protoMsg)
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	test.addedEventIds = test.addedEventIds[:len(test.addedEventIds)-1]
	test.deletedEventsIds = append(test.deletedEventsIds, lastAddedEventId)
	test.responseStatusCode = status.Code(err)

	return nil
}

func (test *eventsTest) iSendGetEventByIdRequestToWithIdOfTheLastCreatedEvent(addr string) error {
	test.setupTestConn(addr)

	client := eventpb.NewEventServiceClient(test.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	lastAddedEventId := test.addedEventIds[len(test.addedEventIds)-1]

	protoMsg := &eventpb.GetEventByIDRequest{Id: lastAddedEventId.String()}

	resp, err := client.GetEventByID(ctx, protoMsg)
	if err != nil {
		return fmt.Errorf("failed to get event: %w", err)
	}

	id, err := uuid.Parse(resp.GetId())
	if err != nil {
		return fmt.Errorf("failed to parse id: %w", err)
	}

	userId, err := uuid.Parse(resp.GetUserId())
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}

	test.eventsReturned = []storage.Event{
		{
			ID:                 id,
			Title:              resp.GetTitle(),
			DateStart:          time.Time{},
			Duration:           storage.PqDuration(resp.GetDuration().AsDuration()),
			Description:        resp.GetDescription(),
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(resp.GetNotificationPeriod().AsDuration()),
			Sent:               resp.GetSent(),
		},
	}
	test.responseStatusCode = status.Code(err)

	return nil
}

func (test *eventsTest) iSendUpdateEventByIdRequestToWithIdOfTheLastCreatedEventAndData(addr string, msg *godog.DocString) error {
	test.setupTestConn(addr)

	client := eventpb.NewEventServiceClient(test.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	protoMsg := &eventpb.UpdateEventByIDRequest{}
	err := jsonpb.Unmarshal(strings.NewReader(msg.Content), protoMsg)
	panicOnErr(err)

	lastAddedEventId := test.addedEventIds[len(test.addedEventIds)-1]
	protoMsg.Id = lastAddedEventId.String()

	_, err = client.UpdateEventByID(ctx, protoMsg)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}

	test.sentData, err = json.Marshal(protoMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal updated event: %w", err)
	}

	test.responseStatusCode = status.Code(err)

	return nil
}

func (test *eventsTest) iSendGetEventsRequestTo(addr string) error {
	test.setupTestConn(addr)

	client := eventpb.NewEventServiceClient(test.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	protoMsg := &eventpb.GetEventsRequest{}

	resp, err := client.GetEvents(ctx, protoMsg)
	if err != nil {
		return fmt.Errorf("failed to get events: %w", err)
	}

	eventsReturned := make([]storage.Event, 0)

	for _, event := range resp.Events {
		id, err := uuid.Parse(event.GetId())
		if err != nil {
			return fmt.Errorf("failed to parse id: %w", err)
		}

		userId, err := uuid.Parse(event.GetUserId())
		if err != nil {
			return fmt.Errorf("failed to parse user id: %w", err)
		}

		eventsReturned = append(eventsReturned, storage.Event{
			ID:                 id,
			Title:              event.GetTitle(),
			DateStart:          time.Time{},
			Duration:           storage.PqDuration(event.GetDuration().AsDuration()),
			Description:        event.GetDescription(),
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(event.GetNotificationPeriod().AsDuration()),
			Sent:               event.GetSent(),
		})
	}

	test.eventsReturned = eventsReturned
	test.responseStatusCode = status.Code(err)

	return nil
}

func (test *eventsTest) iSendGetEventsForTheNextHoursRequestTo(hours int, addr string) error {
	test.setupTestConn(addr)

	client := eventpb.NewEventServiceClient(test.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	protoMsg := &eventpb.GetEventsRequest{
		PeriodStart: timestamppb.New(time.Now().UTC()),
		PeriodEnd:   timestamppb.New(time.Now().Add(time.Duration(hours) * time.Hour).UTC()),
	}

	resp, err := client.GetEvents(ctx, protoMsg)
	if err != nil {
		return fmt.Errorf("failed to get events: %w", err)
	}

	eventsReturned := make([]storage.Event, 0)

	for _, event := range resp.Events {
		id, err := uuid.Parse(event.GetId())
		if err != nil {
			return fmt.Errorf("failed to parse id: %w", err)
		}

		userId, err := uuid.Parse(event.GetUserId())
		if err != nil {
			return fmt.Errorf("failed to parse user id: %w", err)
		}

		eventsReturned = append(eventsReturned, storage.Event{
			ID:                 id,
			Title:              event.GetTitle(),
			DateStart:          time.Time{},
			Duration:           storage.PqDuration(event.GetDuration().AsDuration()),
			Description:        event.GetDescription(),
			UserID:             userId,
			NotificationPeriod: storage.PqDuration(event.GetNotificationPeriod().AsDuration()),
			Sent:               event.GetSent(),
		})
	}

	test.eventsReturned = eventsReturned
	test.responseStatusCode = status.Code(err)

	return nil
}

func (test *eventsTest) itShouldSendNotificationAndMarkCorrespondingEventAsSentWithinSeconds(seconds int) error {
	time.Sleep(time.Duration(seconds) * time.Second)

	lastAddedEventId := test.addedEventIds[len(test.addedEventIds)-1]

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var count int
	err := test.db.GetContext(ctx, &count, `SELECT count(*) FROM events WHERE id = $1 AND sent = true`, lastAddedEventId)
	if err != nil {
		return fmt.Errorf("failed to count events with id %s and sent status equal to true from db: %w", lastAddedEventId, err)
	}

	if count != 1 {
		return fmt.Errorf("failed to send notification for event with id %s", lastAddedEventId)
	}

	return nil
}

func (test *eventsTest) theSenderMicroserviceIsUpAndRunning() error {
	// TODO: Add health check probes to ensure that sender microservice is up and running

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Add expired event that should be immediately processed by sender
	sql := `
		INSERT INTO events (title, date_start, duration, description, user_id, notification_period)
		VALUES (:title, :date_start, :duration, :description, :user_id, :notification_period)
		RETURNING *
	`

	userId, err := uuid.Parse(preparedUserId)
	panicOnErr(err)

	testEvent := storage.Event{
		Title:              "event for notification testing",
		DateStart:          time.Now().Add(-10 * time.Minute).UTC(),
		Duration:           storage.PqDuration(2 * time.Hour),
		Description:        "some description",
		UserID:             userId,
		NotificationPeriod: storage.PqDuration(0),
	}

	row, err := test.db.NamedQueryContext(ctx, sql, testEvent)
	panicOnErr(err)

	var createdEvent storage.Event

	row.Next()
	err = row.StructScan(&createdEvent)
	panicOnErr(err)
	row.Close()

	test.addedEventIds = append(test.addedEventIds, createdEvent.ID)

	return nil
}

func (test *eventsTest) theNumberOfEventsReturnedShouldBe(num int) error {
	receivedNum := len(test.eventsReturned)

	if num != receivedNum {
		return fmt.Errorf("wrong number of events returned: %d expected - %d received", num, receivedNum)
	}

	return nil
}

func (test *eventsTest) theDatabaseShouldContainCreatedEvent() error {
	lastAddedEventId := test.addedEventIds[len(test.addedEventIds)-1]

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var count int
	err := test.db.GetContext(ctx, &count, `SELECT count(*) FROM events WHERE id = $1`, lastAddedEventId)
	if err != nil {
		return fmt.Errorf("failed to count events with id %s from db: %w", lastAddedEventId, err)
	}

	if count != 1 {
		return fmt.Errorf("failed to create event with id %s", lastAddedEventId)
	}

	return nil
}

func (test *eventsTest) theDatabaseShouldContainUpdatedEvent() error {
	lastAddedEventId := test.addedEventIds[len(test.addedEventIds)-1]

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updatedEventFromDb storage.Event
	err := test.db.GetContext(ctx, &updatedEventFromDb, `SELECT * FROM events WHERE id = $1`, lastAddedEventId)
	if err != nil {
		return fmt.Errorf("failed to get event with id %s from db: %w", lastAddedEventId, err)
	}

	var sentEvent eventpb.UpdateEventByIDRequest
	err = json.Unmarshal(test.sentData, &sentEvent)
	if err != nil {
		return fmt.Errorf("failed to unmarshal update event sent data: %w", err)
	}

	if updatedEventFromDb.Title != sentEvent.Title ||
		updatedEventFromDb.Description != sentEvent.Description {
		return fmt.Errorf("failed to update fields for event with id %s", lastAddedEventId)
	}

	return nil
}

func (test *eventsTest) theDatabaseShouldNotContainDeletedEvent() error {
	lastDeletedEventId := test.deletedEventsIds[len(test.deletedEventsIds)-1]

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var count int
	err := test.db.GetContext(ctx, &count, `SELECT count(*) FROM events WHERE id = $1`, lastDeletedEventId)
	if err != nil {
		return fmt.Errorf("failed to count events with id %s from db: %w", lastDeletedEventId, err)
	}

	if count != 0 {
		return fmt.Errorf("failed to delete event with id %s", lastDeletedEventId)
	}

	return nil
}

func (test *eventsTest) theResponseCodeShouldBe(responseCode int) error {
	if test.responseStatusCode != codes.Code(responseCode) {
		return fmt.Errorf("wrong response code: %d", responseCode)
	}

	return nil
}

func InitializeScenario(s *godog.ScenarioContext) {
	test := new(eventsTest)

	s.Before(test.setupTests)

	// Create event
	s.Step(`^I send add event request to "([^"]*)" with data:$`, test.iSendAddEventRequestToWithData)
	s.Step(`^The database should contain created event$`, test.theDatabaseShouldContainCreatedEvent)

	// Update event
	s.Step(`^I send update event by id request to "([^"]*)" with id of the last created event and data:$`, test.iSendUpdateEventByIdRequestToWithIdOfTheLastCreatedEventAndData)
	s.Step(`^The database should contain updated event$`, test.theDatabaseShouldContainUpdatedEvent)

	// Get event
	s.Step(`^I send get event by id request to "([^"]*)" with id of the last created event$`, test.iSendGetEventByIdRequestToWithIdOfTheLastCreatedEvent)

	// Get events
	s.Step(`^I send get events request to "([^"]*)"$`, test.iSendGetEventsRequestTo)
	s.Step(`^I send get events for the next (\d+) hours request to "([^"]*)"$`, test.iSendGetEventsForTheNextHoursRequestTo)
	s.Step(`^The number of events returned should be (\d+)$`, test.theNumberOfEventsReturnedShouldBe)

	// Delete event
	s.Step(`^I send delete event by id request to "([^"]*)" with id of the last created event$`, test.iSendDeleteEventByIdRequestToWithIdOfTheLastCreatedEvent)
	s.Step(`^The database should not contain deleted event$`, test.theDatabaseShouldNotContainDeletedEvent)

	// Send notification
	s.Step(`^The sender microservice is up and running$`, test.theSenderMicroserviceIsUpAndRunning)
	s.Step(`^It should send notification and mark corresponding event as sent within (\d+) seconds$`, test.itShouldSendNotificationAndMarkCorrespondingEventAsSentWithinSeconds)

	// Common
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)

	s.After(test.cleanupTests)
}
