package zipkin

import (
	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	exporter.Set("zipkin", func() trace.SpanExporter {
		exp, err := zipkin.New("")
		if err != nil {
			otel.Handle(err)
			return nil
		}
		return exp
	})
}
