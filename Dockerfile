# Dockerfile for github.com/alpkeskin/mosint

# Build
FROM golang:1.21.0-alpine AS builder
WORKDIR /app
COPY . /app
WORKDIR /app/v3
RUN go mod download
RUN go build ./cmd/mosint

# Release
FROM alpine:3.18.3
COPY --from=builder /app/v3/mosint /usr/local/bin/

# Copy config file ( Change this to your own config file )
COPY --from=builder /app/example-config.yaml /root/.mosint.yaml

ENTRYPOINT ["mosint"]