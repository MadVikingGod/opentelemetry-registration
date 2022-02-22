package otlp

import (
	"context"
	"os"

	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	protocol, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_PROTOCOL")
	if !ok {
		protocol, ok = os.LookupEnv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL")
	}
	var exp *otlptrace.Exporter
	if protocol == "grpc" {
		exp = otlptracegrpc.NewUnstarted()
	} else {
		exp = otlptracehttp.NewUnstarted()
	}

	exporter.Set("otlp", func() trace.SpanExporter {
		err := exp.Start(context.TODO())
		if err != nil {
			otel.Handle(err)
			return nil
		}
		return exp
	})
}
