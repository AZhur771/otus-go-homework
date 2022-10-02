package internalhttp

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Server struct {
	internal *http.Server
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type Application interface { // TODO
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello-world")
}

func NewServer(logger Logger, app Application, host string, port int) *Server {
	mux := http.DefaultServeMux

	mux.Handle("/", loggingMiddleware(http.HandlerFunc(HelloWorldHandler), logger))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}

	return &Server{
		internal: server,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := s.internal.ListenAndServe()
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.internal.Shutdown(ctx)
	return err
}
