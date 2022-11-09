package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Connection interface {
	Channel() (*amqp.Channel, error)
}
