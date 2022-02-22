package all_test

import (
	"fmt"
	"os"

	"github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/all"
)

func Example_jaeger() {
	oldVal := swapEnv("jaeger")
	fmt.Printf("%T\n", exporter.GetExporter())

	swapEnv("otlp")
	fmt.Printf("%T\n", exporter.GetExporter())

	swapEnv("stdout")
	fmt.Printf("%T\n", exporter.GetExporter())

	swapEnv("zipkin")
	fmt.Printf("%T\n", exporter.GetExporter())

	// Output:
	// *jaeger.Exporter
	// *otlptrace.Exporter
	// *stdouttrace.Exporter
	// *zipkin.Exporter

	swapEnv(oldVal)

}

func swapEnv(name string) string {
	old := os.Getenv("OTEL_TRACES_EXPORTER")
	os.Setenv("OTEL_TRACES_EXPORTER", name)
	return old
}
