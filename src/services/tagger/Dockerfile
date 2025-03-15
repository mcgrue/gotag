FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install protoc and required tools
RUN apk add --no-cache protobuf-dev protobuf

# Copy go mod files
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy source code
COPY ./ .

# Install protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate protobuf code and build
RUN protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	protos/tagger/tagger.proto
RUN go build -o /gotag ./server

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /gotag .

EXPOSE 50051
CMD ["./gotag"]