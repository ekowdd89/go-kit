package tracing

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	// semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)



func newConsoleExporter()(oteltrace.SpanExporter, error){
	return stdouttrace.New()
}


func newOtelExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	insecureOpt:= otlptracehttp.WithInsecure()
	endpointOpt:= otlptracehttp.WithEndpoint("otel-collector:4318")

	return otlptracehttp.New(ctx, endpointOpt, insecureOpt)
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider{
	r, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", "go-kit"),
		),
	)
	if err !=nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func New(ctx context.Context)  (tp *sdktrace.TracerProvider ,err error ){

	exp, err := newOtelExporter(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("exporting to %s", exp)
	tp = newTraceProvider(exp)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return
}
