package main

import (
	"fmt"

	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/all"
)

func main() {
	// This only retrieves an exporter
	exp := exporter.GetExporter()
	fmt.Printf("%T\n", exp)

	// You would use `exporter.SetGlobalTracerProvider()` in normal code
}
