package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type exchangeOpts struct {
	name    string
	kind    string
	durable bool
}

type queueOpts struct {
	name       string
	durable    bool
	bindingKey string
}

func setupExchangeAndQueue(ch *amqp.Channel, eOpts exchangeOpts, qOpts queueOpts) error {
	err := ch.ExchangeDeclare(
		eOpts.name,
		eOpts.kind,
		eOpts.durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error while declaring exchange: %w", err)
	}

	queue, err := ch.QueueDeclare(
		qOpts.name,
		qOpts.durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error while declaring queue: %w", err)
	}

	err = ch.QueueBind(
		queue.Name,
		qOpts.bindingKey,
		eOpts.name,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error while binding queue to exchange: %w", err)
	}

	return nil
}
