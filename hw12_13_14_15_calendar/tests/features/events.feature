Feature: Event crud grpc api
  As a user of event grpc api
  I should be able to perform all the crud actions

  Scenario: Create event
    When I send add event request to "grpc://calendar:3000" with data:
    """
    {
      "user_id": "19a12b49-a57a-4f1e-8e66-152be08e6165",
      "title": "some_title",
      "description": "some_description",
      "date_start": "2020-05-22T20:32:05Z",
      "duration": "2h",
      "notification_period": "2h"
    }
    """
    Then The response code should be 0
    And The database should contain created event

    When I send add event request to "grpc://calendar:3000" with data:
    """
    {
      "user_id": "19a12b49-a57a-4f1e-8e66-152be08e6165",
      "title": "some_title2",
      "description": "some_description2",
      "date_start": "2021-05-22T20:32:05Z",
      "duration": "2h",
      "notification_period": "2h"
    }
    """
    Then The response code should be 0
    And The database should contain created event

  Scenario: Update event
    When I send update event by id request to "grpc://calendar:3000" with id of the last created event and data:
    """
    {
      "user_id": "19a12b49-a57a-4f1e-8e66-152be08e6165",
      "title": "another_title",
      "description": "another_description",
      "date_start": "2020-05-22T20:32:05Z",
      "duration": "2h",
      "notification_period": "2h"
    }
    """
    Then The response code should be 0
    And The database should contain updated event

  Scenario: Get event
    When I send get event by id request to "grpc://calendar:3000" with id of the last created event
    Then The response code should be 0

  Scenario: Get events
    When I send get events request to "grpc://calendar:3000"
    Then The response code should be 0
    And The number of events returned should be 4

    When I send get events for the next 1 hours request to "grpc://calendar:3000"
    Then The response code should be 0
    And The number of events returned should be 2

    When I send get events for the next 2 hours request to "grpc://calendar:3000"
    Then The response code should be 0
    And The number of events returned should be 3

  Scenario: Delete event
    When I send delete event by id request to "grpc://calendar:3000" with id of the last created event
    Then The response code should be 0
    And The database should not contain deleted event

  Scenario: Send notification
    When The sender microservice is up and running
    Then It should send notification and mark corresponding event as sent within 60 seconds
