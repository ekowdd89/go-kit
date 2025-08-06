package otel

import (
	"context"
	"log"

	// "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func InitTracer(ctx context.Context) func() {
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("tempo:4318"), // OTLP HTTP endpoint dari Tempo
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("my-go-service"),
		)),
	)
	otel.SetTracerProvider(tp)
	return func() {
		_ = tp.Shutdown(ctx)
	}
}
