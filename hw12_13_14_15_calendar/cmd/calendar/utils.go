package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	inmemorystorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/storage/sql"
	toml "github.com/pelletier/go-toml"
)

func getStorage(dbConf DatabaseConf) (app.Storage, error) {
	if dbConf.MemoryStorage == "sql" {
		storage := sqlstorage.New()
		ctx := context.Background()
		datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.DBName, dbConf.SslMode)
		err := storage.Connect(ctx, datasource, dbConf.MaxConnections)
		return storage, err
	} else if dbConf.MemoryStorage == "inmemory" {
		return inmemorystorage.New(), nil
	}

	log.Fatalf(fmt.Sprintf("unknown storage type: %s", dbConf.MemoryStorage))
	return nil, nil
}

func getConfig(configFile string) (Config, error) {
	config := NewConfig()

	f, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("error while open config file %s: %v\n", configFile, err)
		return config, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("error while read config file %s: %v\n", configFile, err)
		return config, err
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("error while unmarshal config file %s: %v\n", configFile, err)
		return config, err
	}

	return config, nil
}
