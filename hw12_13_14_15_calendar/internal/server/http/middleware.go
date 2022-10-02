package internalhttp

import (
	"fmt"
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/internal/app"
	"net/http"
	"strings"
	"time"
)

type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

func loggingMiddleware(next http.Handler, logg app.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestAccepted := time.Now()
		o := &responseObserver{ResponseWriter: w}
		next.ServeHTTP(o, r)
		latency := time.Now().Sub(requestAccepted)
		addr := r.RemoteAddr
		if i := strings.LastIndex(addr, ":"); i != -1 {
			addr = addr[:i]
		}
		logg.Info(fmt.Sprintf(
			"%s - %s - %s - %s - %d - %s - %s",
			addr, r.Method, r.RequestURI, r.Proto, o.status, latency, r.UserAgent()))
	})
}
