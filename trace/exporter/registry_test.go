package exporter

import (
	"os"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/sdk/trace"
)

func TestGetExporter(t *testing.T) {
	tests := []struct {
		name    string
		envName string
		want    trace.SpanExporter
	}{
		{
			name:    "Exporter doesnot exist",
			envName: "doesNotExist",
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("OTEL_TRACES_EXPORTER", tt.envName)
			if got := GetExporter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
