package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/server/http"
	inmemorystorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/sql"
	_ "github.com/lib/pq"
	toml "github.com/pelletier/go-toml"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

func getStorage(dbConf DatabaseConf) (app.Storage, error) {
	if dbConf.InMemoryStorage {
		return inmemorystorage.New(), nil
	}

	storage := sqlstorage.New()
	ctx := context.Background()
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.DBName, dbConf.SslMode)
	err := storage.Connect(ctx, datasource)
	return storage, err
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig()

	f, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("error while open config file %s: %v\n", configFile, err)
		return
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("error while read config file %s: %v\n", configFile, err)
		return
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("error while unmarshal config file %s: %v\n", configFile, err)
		return
	}

	storage, err := getStorage(config.Database)
	if err != nil {
		fmt.Printf("error while getting storage: %v\n", err)
		return
	}

	logg := logger.New(config.Logger.Level)

	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar, config.Server.Host, config.Server.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), config.Server.ShutdownTimeout)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
