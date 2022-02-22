package stdout

import (
	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	exporter.Set("stdout", func() trace.SpanExporter {
		exp, err := stdouttrace.New()
		if err != nil {
			otel.Handle(err)
			return nil
		}
		return exp
	})
}
