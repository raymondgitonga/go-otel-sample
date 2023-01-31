package httpserver

import (
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

const name = "otel-sample"

type Handler struct {
	tracer trace.Tracer
}

func NewHandler(tracer trace.Tracer) *Handler {
	return &Handler{
		tracer: tracer,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, span := h.tracer.Start(ctx, "health-check-endpoint")
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
