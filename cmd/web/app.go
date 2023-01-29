package web

import (
	"github.com/raymondgitonga/go-otel-sample/internal/adapters/httpserver"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	tracer trace.Tracer
	logger *logrus.Logger
}

func NewApp(tracer trace.Tracer, logger *logrus.Logger) *App {
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
