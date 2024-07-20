# Builder stage
FROM golang:1.22.5-alpine3.19 AS builder

# Create and change to the build directoy
WORKDIR /build

# Copy from all files in the grpc_server to container
COPY . .

# Download dependencies from container
RUN go mod download

# Build inside container
RUN go build -o ./basic_grpc_server ./cmd/main.go

# create docker image for basic_grpc_server based on alpine linux
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

# set env port listener
ENV BASIC_GRPC_PORT=8099
EXPOSE 8099

# copy basic_grpc_server binary from builder stage
COPY --from=builder /build/basic_grpc_server ./basic_grpc_server

CMD ["/basic_grpc_server"]