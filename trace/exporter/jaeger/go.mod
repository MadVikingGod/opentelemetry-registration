module github.com/madvikinggod/opentelemetry-registration/trace/exporter/jaeger

go 1.17

require (
	github.com/madvikinggod/opentelemetry-registration/trace/exporter v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/exporters/jaeger v1.4.1
	go.opentelemetry.io/otel/sdk v1.4.1
)

require (
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/trace v1.4.1 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
)

replace github.com/madvikinggod/opentelemetry-registration/trace/exporter => ../
