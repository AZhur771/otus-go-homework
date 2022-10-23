package internalhttp

import (
	"context"
	"fmt"
	eventpb "github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/api/stubs"

	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"

	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	grpcGW *http.Server
}

// NewServer returns new grpc Gateway Server
func NewServer(ctx context.Context, logger app.Logger, host string, port int, conn *grpc.ClientConn) (*Server, error) {
	gwmux := runtime.NewServeMux()

	err := eventpb.RegisterEventServiceHandler(ctx, gwmux, conn)
	if err != nil {
		return nil, err
	}

	oa, err := GetOpenAPIHandler()
	if err != nil {
		return nil, err
	}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			gwmux.ServeHTTP(w, r)
			return
		}
		oa.ServeHTTP(w, r)
	})

	handler := loggingMiddleware(handlerFunc, logger)

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: handler,
	}

	return &Server{
		grpcGW: gwServer,
	}, nil
}

func (s *Server) Start() error {
	err := s.grpcGW.ListenAndServe()
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.grpcGW.Shutdown(ctx)
	return err
}
