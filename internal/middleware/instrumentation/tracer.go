package instrumentation

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	endpoint = "collector:4318"
)

var (
	ServiceName    = "go-otel-sample"
	ServiceVersion = "0.1.0"
)

func SetupTracer(ctx context.Context) (func() error, trace.Tracer) {
	tp, err := newTraceProvider(ctx)

	if err != nil {
		fmt.Println("error setting up trace Provider ", err)
		return nil, nil
	}

	return func() error {
		return tp.Shutdown(ctx)
	}, tp.Tracer("main-tracer")
}

func newTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	r, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error setting up resources, %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	traceExporter, err := setupTraceExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("error seting up trace exportet, %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(r),
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(traceExporter)),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider, nil
}

func setupTraceExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	conn, err := grpc.DialContext(ctx, "localhost:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server, %w", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("error initialising trace exporter, %w", err)
	}

	return traceExporter, nil
}
