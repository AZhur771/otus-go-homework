package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/config"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/sender"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config_sender.toml", "Path to sender configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	senderConfig, err := config.GetSenderConfig(configFile)
	if err != nil {
		log.Fatalf("error while getting config: %s", err)
	}

	conn, err := amqp.Dial(config.GetAMQPURI(senderConfig.RabbitMQ))
	if err != nil {
		log.Fatalf("error while dialing RabbitMQ: %s", err)
	}

	logg := logger.New(senderConfig.Logger.Level)

	consumer, err := config.GetConsumer(conn, logg)
	if err != nil {
		log.Fatalf("error while getting message producer: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	err = consumer.Connect(
		ctx,
		senderConfig.RabbitMQ.Exchange,
		senderConfig.RabbitMQ.ExchangeType,
		senderConfig.RabbitMQ.Queue,
		senderConfig.RabbitMQ.ConsumerTag,
		senderConfig.RabbitMQ.BindingKey,
		senderConfig.RabbitMQ.Persistent,
	)
	if err != nil {
		log.Panicf("error while setting up exchange and queue: %s", err)
	}

	go func() {
		<-ctx.Done()
		if err = consumer.Disconnect(); err != nil {
			logg.Error(fmt.Sprintf("error while closing channel: %s", err))
		}
		os.Exit(1)
	}()

	sender := sender.New(logg, consumer)

	logg.Info("sender is up and running")

	if err := sender.Run(ctx); err != nil {
		logg.Error(fmt.Sprintf("sender failed: %s", err))
		cancel()
	}
}
