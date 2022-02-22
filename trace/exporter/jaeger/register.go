package jaeger

import (
	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	exporter.Set("jaeger", func() trace.SpanExporter {
		exp, err := jaeger.New(jaeger.WithAgentEndpoint())
		if err != nil {
			otel.Handle(err)
			return nil
		}
		return exp
	})
}
