module pingpawn.com/gotag

go 1.22.0

require (
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.4
)

require (
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)

replace pingpawn.com/gotag/protos/helloworld => ./protos

// replace google.golang.org/grpc => ../

// replace google.golang.org/grpc/gcp/observability => ../gcp/observability

// replace google.golang.org/grpc/stats/opentelemetry => ../stats/opentelemetry
