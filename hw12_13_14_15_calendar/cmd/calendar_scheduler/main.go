package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/config"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/scheduler"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config_scheduler.toml", "Path to scheduler configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	schedulerConfig, err := config.GetSchedulerConfig(configFile)
	if err != nil {
		log.Fatalf("error while getting config: %s", err)
	}

	if schedulerConfig.Database.StorageType != "sql" {
		log.Fatal("for scheduler only sql database is supported")
	}

	storage, err := config.GetStorage(schedulerConfig.Database)
	if err != nil {
		log.Fatalf("error while getting storage: %s", err)
	}

	logg := logger.New(schedulerConfig.Logger.Level)

	conn, err := amqp.Dial(config.GetAMQPURI(schedulerConfig.RabbitMQ))
	if err != nil {
		log.Fatalf("error while dialing RabbitMQ: %s", err)
	}

	producer, err := config.GetProducer(conn, logg)
	if err != nil {
		log.Fatalf("error while getting message producer: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	err = producer.Connect(
		ctx,
		schedulerConfig.RabbitMQ.Exchange,
		schedulerConfig.RabbitMQ.ExchangeType,
		schedulerConfig.RabbitMQ.Queue,
		schedulerConfig.RabbitMQ.BindingKey,
		schedulerConfig.RabbitMQ.Persistent,
		schedulerConfig.RabbitMQ.Reliable,
	)
	if err != nil {
		log.Panicf("error while setting up exchange and queue: %s", err)
	}

	go func() {
		if err = producer.WaitForConfirms(ctx); err != nil {
			logg.Error(fmt.Sprintf("error while setting channel in confirm mode: %s", err))
		}
	}()

	go func() {
		<-ctx.Done()
		if err = producer.Disconnect(); err != nil {
			logg.Error(fmt.Sprintf("error while closing channel: %s", err))
		}
		os.Exit(1)
	}()

	schedulr := scheduler.New(
		storage,
		logg,
		producer,
		time.Duration(schedulerConfig.Scheduler.ScanPeriod)*time.Minute,
		time.Duration(schedulerConfig.Scheduler.DeletePeriod)*time.Minute,
		schedulerConfig.Scheduler.StartImmediately,
	)

	logg.Info("scheduler is up and running")

	go func() {
		if err = schedulr.RunDeleter(ctx); err != nil {
			logg.Error(fmt.Sprintf("scheduler failed to run deleter: %s", err))
			cancel()
		}
	}()

	if err = schedulr.RunNotifier(ctx); err != nil {
		logg.Error(fmt.Sprintf("scheduler failed to run notifier: %s", err))
		cancel()
	}
}
