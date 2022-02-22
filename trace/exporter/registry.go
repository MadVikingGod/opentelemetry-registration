package exporter

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

var _registry = sync.Map{}

func Set(name string, exporterProvider func() trace.SpanExporter) {
	_registry.Store(strings.ToLower(name), exporterProvider)
}

func GetExporter() trace.SpanExporter {
	name, ok := os.LookupEnv("OTEL_TRACES_EXPORTER")
	if !ok || name == "" {
		name = "otlp"
	}

	exp, ok := _registry.Load(strings.ToLower(name))
	if !ok {
		otel.Handle(fmt.Errorf("Could not load exporter \"%s\". It was not registered.", name))
		return nil
	}
	return exp.(func() trace.SpanExporter)()
}

func GetBatchSpanProcessor() trace.SpanProcessor {
	return trace.NewBatchSpanProcessor(GetExporter())
}

// Do not call GetTracerProvider in `init()`, as anything registering may be run in a different order.
func GetTracerProvider() *trace.TracerProvider {
	return trace.NewTracerProvider(trace.WithBatcher(GetExporter()))
}

func SetGlobalTracerProvider() {
	otel.SetTracerProvider(GetTracerProvider())
}
