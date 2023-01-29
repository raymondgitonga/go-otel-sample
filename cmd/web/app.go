package web

import (
	"github.com/raymondgitonga/go-otel-sample/internal/adapters/httpserver"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	tracer trace.Tracer
	logger *zap.Logger
}

func NewApp(tracer trace.Tracer, logger *zap.Logger) *App {
	return &App{
		tracer: tracer,
		logger: logger,
	}
}

func (c *App) StartApp() (*mux.Router, error) {
	r := mux.NewRouter()
	handler := httpserver.NewHandler(c.tracer, c.logger)

	r.HandleFunc("/health-check", handler.HealthCheck).Methods(http.MethodGet)

	return r, nil
}
