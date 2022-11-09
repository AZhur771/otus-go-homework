package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	eventpb "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/api/stubs"
	config "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/config"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/server/http"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	calendarConfig, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatalf("error while getting config: %s", err)
	}

	calendarStorage, err := config.GetStorage(calendarConfig.Database)
	if err != nil {
		log.Fatalf("error while getting storage: %s", err)
	}

	logg := logger.New(calendarConfig.Logger.Level)

	grpcClientConn, err := internalgrpc.NewClientConn(
		context.Background(),
		calendarConfig.Server.Host,
		calendarConfig.Server.Port,
	)
	if err != nil {
		log.Fatalf("error while instantiating grpc client connection: %s", err)
	}

	grpcGWServer, err := internalhttp.NewServer(
		context.Background(),
		logg,
		calendarConfig.Server.Host,
		calendarConfig.Server.GatewayPort,
		grpcClientConn,
	)
	if err != nil {
		log.Fatalf("error while instantiating grpc gateway server: %s", err)
	}

	grpcServer := grpc.NewServer()
	eventpb.RegisterEventServiceServer(grpcServer, internalgrpc.NewServer(calendarStorage, logg))

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		if err := grpcGWServer.Start(); err != nil {
			logg.Error("failed to start grpc gateway server: " + err.Error())
			cancel()
		}
	}()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(calendarConfig.Server.ShutdownTimeout)*time.Millisecond,
		)
		defer cancel()

		if err := grpcGWServer.Stop(ctx); err != nil {
			logg.Error("failed to stop grpc gateway server gracefully: " + err.Error())
		}

		grpcServer.GracefulStop()
		os.Exit(1) //nolint:gocritic,nolintlint
	}()

	logg.Info("grpc server is up and running")
	logg.Info(fmt.Sprintf("http server is up and running at http://%s:%d",
		calendarConfig.Server.Host, calendarConfig.Server.GatewayPort))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", calendarConfig.Server.Host, calendarConfig.Server.Port))
	if err != nil {
		logg.Error("failed to get tcp listener: " + err.Error())
		cancel()
	}

	if err := grpcServer.Serve(lis); err != nil {
		logg.Error("failed to start grpc server: " + err.Error())
		cancel()
	}
}
