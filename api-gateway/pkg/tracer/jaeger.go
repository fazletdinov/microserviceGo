package tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer(ctx context.Context, jaegerURL string, serviceName string) (*sdktrace.TracerProvider, error) {
	exporter, err := NewJaegerExporter(ctx, jaegerURL)
	if err != nil {
		return nil, fmt.Errorf("initialize exporter: %w", err)
	}

	tp, err := NewTracerProvider(exporter, serviceName)
	if err != nil {
		return nil, fmt.Errorf("initialize provider: %w", err)
	}

	return tp, nil
}

func NewJaegerExporter(ctx context.Context, jaegerURL string) (*otlptrace.Exporter, error) {
	traceExporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(jaegerURL),
		),
	)
	if err != nil {
		return nil, err
	}
	fmt.Printf("traceExpoter ================== %+v\n", traceExporter)
	return traceExporter, nil
}

func NewTracerProvider(traceExporter *otlptrace.Exporter, serviceName string) (*sdktrace.TracerProvider, error) {
	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.DeploymentEnvironmentKey.String("develop"),
		),
	)
	if err != nil {
		return nil, err
	}
	fmt.Printf("resourse ================== %+v\n", resource)

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(traceExporter)),
		sdktrace.WithSyncer(traceExporter),
		sdktrace.WithResource(resource),
	)

	fmt.Printf("provider ================== %+v\n", provider)

	otel.SetTracerProvider(provider)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return provider, nil
}
