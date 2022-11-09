package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/streadway/amqp"
)

type Consumer struct {
	queueName   string
	consumerTag string
	ch          *amqp.Channel
	logger      app.Logger
}

func NewConsumer(conn Connection, logger app.Logger) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	return &Consumer{
		ch:     ch,
		logger: logger,
	}, nil
}

func (c *Consumer) Connect(
	ctx context.Context,
	exchangeName, exchangeType, queueName, consumerTag, bindingKey string,
	persistent bool,
) error {
	err := setupExchangeAndQueue(
		c.ch,
		exchangeOpts{
			name:    exchangeName,
			kind:    exchangeType,
			durable: persistent,
		},
		queueOpts{
			name:       queueName,
			durable:    persistent,
			bindingKey: bindingKey,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to setup exchange and queue: %w", err)
	}

	c.queueName = queueName
	c.consumerTag = consumerTag

	return nil
}

func (c *Consumer) Consume(ctx context.Context) (<-chan app.Message, error) {
	messages := make(chan app.Message)

	deliveries, err := c.ch.Consume(c.queueName, c.consumerTag, false, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("error while starting to consumer messages: %w", err)
	}

	go func() {
		defer func() {
			close(messages)
			c.logger.Info(fmt.Sprintf("stopping consumer %s", c.consumerTag))
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case del := <-deliveries:
				if err := del.Ack(false); err != nil {
					c.logger.Error(fmt.Sprintf("error while acknowledging message: %s", err))
				}

				msg := &app.Message{}
				err := json.Unmarshal(del.Body, msg)
				if err != nil {
					c.logger.Error(fmt.Sprintf("error while unmarshalling message: %s", err))
					continue
				}

				select {
				case <-ctx.Done():
					return
				case messages <- *msg:
				}
			}
		}
	}()

	return messages, nil
}

func (c *Consumer) Disconnect() error {
	if err := c.ch.Close(); err != nil {
		return fmt.Errorf("error while closing channel: %w", err)
	}

	return nil
}
