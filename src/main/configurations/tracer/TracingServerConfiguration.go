package main_configurations_tracer

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_utils "mpindicator/main/utils"
)

const MSG_ERROR_INITIATING_EXPORTER = "Error initiating zipkin exporter"

const IDX_TRACING_SERVICE_NAME_KEY = "Tracing.service.name.key"
const IDX_TRACING_SERVER_URL = "Tracing.server.url"

func InitTracer() {
	tracingServiceNameKey := main_configurations_yml.GetBeanPropertyByName(IDX_TRACING_SERVICE_NAME_KEY)
	tracingServiceUrl := main_configurations_yml.GetBeanPropertyByName(IDX_TRACING_SERVER_URL)

	exporter, err := zipkin.New(
		tracingServiceUrl,
	)
	main_utils.FailOnError(err, MSG_ERROR_INITIATING_EXPORTER)

	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(tracingServiceNameKey),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, propagation.Baggage{}))
}
