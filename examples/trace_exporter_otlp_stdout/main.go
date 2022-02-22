package main

import (
	"fmt"

	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/otlp"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/stdout"
)

func main() {
	// This only retrieves an exporter
	exp := exporter.GetExporter()
	fmt.Printf("%T\n", exp)

	// You would use `exporter.SetGlobalTracerProvider()` in normal code
}
