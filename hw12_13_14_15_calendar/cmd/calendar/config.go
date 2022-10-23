package main

import (
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
	"time"
)

// Config При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger   LoggerConf
	Server   ServerConf
	Database DatabaseConf
}

type LoggerConf struct {
	Level logger.Level
}

type ServerConf struct {
	Host            string        `toml:"host"`
	Port            int           `toml:"port"`
	GatewayPort     int           `toml:"gateway_port"`
	ShutdownTimeout time.Duration `toml:"shutdown_timeout"`
}

type DatabaseConf struct {
	MemoryStorage  string        `toml:"memory_storage"`
	DBTimeout      time.Duration `toml:"db_timeout"`
	MaxConnections int           `toml:"max_connections"`
	Host           string        `toml:"host"`
	Port           int           `toml:"port"`
	Username       string        `toml:"username"`
	Password       string        `toml:"password"`
	DBName         string        `toml:"db_name"`
	SslMode        string        `toml:"ssl_mode"`
}

func NewConfig() Config {
	return Config{}
}
