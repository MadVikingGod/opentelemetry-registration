#! /usr/bin/env bash

for exp in otlp stdout jaeger zipkin doesNotExist; do
    echo "Starting example with $exp exporter"
    OTEL_TRACES_EXPORTER=$exp go run .
done