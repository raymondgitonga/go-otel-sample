package web

import (
	"github.com/raymondgitonga/go-otel-sample/internal/adapters/httpserver"
	"go.opentelemetry.io/otel/trace"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	tracer trace.Tracer
}

func NewApp(tracer trace.Tracer) *App {
	return &App{
		tracer: tracer,
	}
}

func (c *App) StartApp() (*mux.Router, error) {
	r := mux.NewRouter()
	handler := httpserver.NewHandler(c.tracer)

	r.HandleFunc("/health-check", handler.HealthCheck).Methods(http.MethodGet)

	return r, nil
}
