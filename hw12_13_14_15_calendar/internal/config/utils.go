package config

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/rabbitmq"
	inmemorystorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pelletier/go-toml"
	"github.com/streadway/amqp"
)

func ReadConfig(configFile string) ([]byte, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("error while open config file %s: %w", configFile, err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error while read config file %s: %w", configFile, err)
	}

	return b, nil
}

func GetConfig(configFile string) (Config, error) {
	config := NewConfig()

	b, err := ReadConfig(configFile)
	if err != nil {
		return config, fmt.Errorf("error while reading config file %s: %w", configFile, err)
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		return config, fmt.Errorf("error while unmarshal config file %s: %w", configFile, err)
	}

	return config, nil
}

func GetSchedulerConfig(configFile string) (SchedulerConfig, error) {
	config := NewSchedulerConfig()

	b, err := ReadConfig(configFile)
	if err != nil {
		return config, fmt.Errorf("error while reading config file %s: %w", configFile, err)
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		return config, fmt.Errorf("error while unmarshal config file %s: %w", configFile, err)
	}

	return config, nil
}

func GetSenderConfig(configFile string) (SenderConfig, error) {
	config := NewSenderConfig()

	b, err := ReadConfig(configFile)
	if err != nil {
		return config, fmt.Errorf("error while reading config file %s: %w", configFile, err)
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		return config, fmt.Errorf("error while unmarshal config file %s: %w", configFile, err)
	}

	return config, nil
}

func GetAMQPURI(rabbitMQConf RabbitMQConf) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d",
		rabbitMQConf.Username, rabbitMQConf.Password, rabbitMQConf.Host, rabbitMQConf.Port)
}

func GetStorage(dbConf DatabaseConf) (app.Storage, error) {
	if dbConf.StorageType == "sql" {
		storage := sqlstorage.New(time.Duration(dbConf.DBTimeout) * time.Millisecond)
		ctx := context.Background()
		datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.DBName, dbConf.SslMode)
		err := storage.Connect(ctx, datasource, dbConf.MaxConnections)
		return storage, err
	} else if dbConf.StorageType == "inmemory" {
		return inmemorystorage.New(), nil
	}

	return nil, fmt.Errorf("unknown storage type: %s", dbConf.StorageType)
}

func GetProducer(conn *amqp.Connection, logger app.Logger) (app.Producer, error) {
	p, err := rabbitmq.NewProducer(conn, logger)
	if err != nil {
		return nil, fmt.Errorf("error while creating producer: %w", err)
	}

	return p, nil
}

func GetConsumer(conn *amqp.Connection, logger app.Logger) (app.Consumer, error) {
	c, err := rabbitmq.NewConsumer(conn, logger)
	if err != nil {
		return nil, fmt.Errorf("error while creating consumer: %w", err)
	}

	return c, nil
}
