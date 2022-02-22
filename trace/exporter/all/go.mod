module github.com/madvikinggod/opentelemetry-registration/trace/exporter/all

go 1.17

replace github.com/madvikinggod/opentelemetry-registration/trace/exporter => ../

replace github.com/madvikinggod/opentelemetry-registration/trace/exporter/jaeger => ../jaeger

replace github.com/madvikinggod/opentelemetry-registration/trace/exporter/otlp => ../otlp

replace github.com/madvikinggod/opentelemetry-registration/trace/exporter/stdout => ../stdout

replace github.com/madvikinggod/opentelemetry-registration/trace/exporter/zipkin => ../zipkin

require (
	github.com/madvikinggod/opentelemetry-registration/trace/exporter/jaeger v0.0.0-00010101000000-000000000000
	github.com/madvikinggod/opentelemetry-registration/trace/exporter/otlp v0.0.0-00010101000000-000000000000
	github.com/madvikinggod/opentelemetry-registration/trace/exporter/stdout v0.0.0-00010101000000-000000000000
	github.com/madvikinggod/opentelemetry-registration/trace/exporter/zipkin v0.0.0-00010101000000-000000000000
)

require (
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/madvikinggod/opentelemetry-registration/trace/exporter v0.0.0-00010101000000-000000000000 // indirect
	github.com/openzipkin/zipkin-go v0.4.0 // indirect
	go.opentelemetry.io/otel v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.4.1 // indirect
	go.opentelemetry.io/otel/sdk v1.4.1 // indirect
	go.opentelemetry.io/otel/trace v1.4.1 // indirect
	go.opentelemetry.io/proto/otlp v0.12.0 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
