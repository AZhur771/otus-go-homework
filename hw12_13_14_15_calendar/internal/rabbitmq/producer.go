package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/streadway/amqp"
)

type Producer struct {
	exchangeName string
	persistent   bool
	reliable     bool
	routingKey   string
	ch           *amqp.Channel
	logger       app.Logger
}

func NewProducer(conn Connection, logger app.Logger) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	return &Producer{
		ch:     ch,
		logger: logger,
	}, nil
}

func (p *Producer) WaitForConfirms(ctx context.Context) error {
	if p.reliable {
		if err := p.ch.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %w", err)
		}

		confirms := p.ch.NotifyPublish(make(chan amqp.Confirmation, 1))

		for {
			select {
			case confirm := <-confirms:
				if confirm.Ack {
					p.logger.Info(fmt.Sprintf("confirmed delivery with delivery tag: %d", confirm.DeliveryTag))
				} else {
					p.logger.Error(fmt.Sprintf("failed delivery of delivery tag: %d", confirm.DeliveryTag))
				}
			case <-ctx.Done():
				return nil
			}
		}
	}

	return nil
}

func (p *Producer) Connect(
	ctx context.Context,
	exchangeName, exchangeType, queueName, bindingKey string,
	persistent, reliable bool,
) error {
	err := setupExchangeAndQueue(
		p.ch,
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

	p.reliable = reliable
	p.persistent = persistent
	p.routingKey = bindingKey
	p.exchangeName = exchangeName

	return nil
}

func (p *Producer) Publish(ctx context.Context, msg app.Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error while marshalling message: %w", err)
	}

	// 1=non-persistent, 2=persistent
	var deliveryMode uint8

	if p.persistent {
		deliveryMode = amqp.Persistent
	} else {
		deliveryMode = amqp.Transient
	}

	if err = p.ch.Publish(
		p.exchangeName, // publish to an exchange
		p.routingKey,   // routing to 0 or more queues
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			Body:            body,
			DeliveryMode:    deliveryMode,
			Priority:        0, // 0-9
		},
	); err != nil {
		return fmt.Errorf("error while publishing message: %w", err)
	}

	return nil
}

func (p *Producer) Disconnect() error {
	if err := p.ch.Close(); err != nil {
		return fmt.Errorf("error while closing channel: %w", err)
	}

	return nil
}
