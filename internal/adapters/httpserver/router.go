package httpserver

import (
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
)

const name = "otel-sample"

type Handler struct {
	tracer trace.Tracer
	logger *zap.Logger
}

func NewHandler(tracer trace.Tracer, logger *zap.Logger) *Handler {
	return &Handler{
		tracer: tracer,
		logger: logger,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, span := h.tracer.Start(r.Context(), "health-check")
	h.logger.Info("health check called")
	response, err := json.Marshal("Healthy")

	if err != nil {
		fmt.Printf("error writing marshalling response: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		fmt.Printf("error writing httpserver response: %s", err)
	}

	span.End()
}
