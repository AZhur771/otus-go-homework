package config

import (
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
)

// Config При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger   LoggerConf   `toml:"logger"`
	Server   ServerConf   `toml:"server"`
	Database DatabaseConf `toml:"database"`
}

type SchedulerConfig struct {
	Logger    LoggerConf    `toml:"logger"`
	Scheduler SchedulerConf `toml:"scheduler"`
	Database  DatabaseConf  `toml:"database"`
	RabbitMQ  RabbitMQConf  `toml:"rabbitmq"`
}

type SenderConfig struct {
	Logger   LoggerConf   `toml:"logger"`
	Database DatabaseConf `toml:"database"`
	RabbitMQ RabbitMQConf `toml:"rabbitmq"`
}

type LoggerConf struct {
	Level logger.Level `toml:"level"`
}

type ServerConf struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	GatewayPort     int    `toml:"gateway_port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

type SchedulerConf struct {
	ScanPeriod       int  `toml:"scan_period"`
	DeletePeriod     int  `toml:"delete_period"`
	StartImmediately bool `toml:"start_immediate"`
}

type RabbitMQConf struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Exchange     string `toml:"exchange"`
	ExchangeType string `toml:"exchange_type"`
	Queue        string `toml:"queue"`
	ConsumerTag  string `toml:"consumer_tag"`
	BindingKey   string `toml:"binding_key"`
	Reliable     bool   `toml:"reliable"`
	Persistent   bool   `toml:"persistent"`
}

type DatabaseConf struct {
	StorageType    string `toml:"storage_type"`
	DBTimeout      int    `toml:"db_timeout"`
	MaxConnections int    `toml:"max_connections"`
	Host           string `toml:"host"`
	Port           int    `toml:"port"`
	Username       string `toml:"username"`
	Password       string `toml:"password"`
	DBName         string `toml:"db_name"`
	SslMode        string `toml:"ssl_mode"`
}

func NewConfig() Config {
	return Config{}
}

func NewSchedulerConfig() SchedulerConfig {
	return SchedulerConfig{}
}

func NewSenderConfig() SenderConfig {
	return SenderConfig{}
}
