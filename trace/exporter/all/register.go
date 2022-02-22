package all

import (
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/jaeger"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/otlp"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/stdout"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/zipkin"
)
