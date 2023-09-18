module github.com/pmarques/ifconfig.me

go 1.16

require (
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.44.0
	go.opentelemetry.io/otel v1.18.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.17.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.17.0
	go.opentelemetry.io/otel/sdk v1.18.0
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
