# Registration Mechanism for Opentelemetry Components in go

This is repository for exploring how we can register different compoents of the opentelemetry project to enable choice via environment variables.  In general the Opentelemetry go SDK does not import any particular exporter or propagator, so if you want to use one you MUST import it into your code, and configure it. For most applications this is a choice made by the application writer, but the API and the SDK do not provide a way to set this.  This library provides a mechanism to have and set this logic, but allow other exporters (not in the mainsdk) to also be registered and used.

## How to use in an application

For an application you first must import the components that you would like to use.

```go
import (
    "github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/otlp"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/stdout"
)
```

or import a collection like the `all` package which imports all from the SDK.

```go
import (
    "github.com/madvikinggod/opentelemetry-registration/trace/exporter"
	_ "github.com/madvikinggod/opentelemetry-registration/trace/exporter/all"
)
```

Then to set your tracer provider you can use `exporter.SetTracerProvider()`.  Remember that when you fetch any of the Exporters, BatchSpanProcessors, or TracerProviders that you will also start the exporters, so you should not call these functions more then once.

Do not call any of the retrieving functions in an `init()` block.  Because the order of init functions in packages is not guaranteed the registration may not happen before the application `init()`.

## How to register an exporter

If you have written an exporter and would like to enable applications to use it in a similar fashion, you need to register you exporter!

It is as simple as including a `func init()` that calls `exporter.Set()`.  this is from the zipkin exporter:

```go
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
```

Use a name that will uniquely define how users will access your exporter, and a function that returns a started exporter.  Then make sure that your exporter has well documented environment variables for setting things like the endpoint, and anything else that might need configuring.