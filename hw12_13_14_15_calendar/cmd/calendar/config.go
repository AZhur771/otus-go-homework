package main

import (
	"time"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
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
	ShutdownTimeout time.Duration `toml:"shutdown_timeout"`
}

type DatabaseConf struct {
	InMemoryStorage bool   `toml:"inmemory_storage"`
	DBTimeout       int    `toml:"db_timeout"`
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	Username        string `toml:"username"`
	Password        string `toml:"password"`
	DBName          string `toml:"db_name"`
	SslMode         string `toml:"ssl_mode"`
}

func NewConfig() Config {
	return Config{}
}
