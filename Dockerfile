FROM golang:1.19.2-alpine AS builder
RUN apk add --no-cache git
RUN go install -v github.com/alpkeskin/mosint@latest

ENTRYPOINT ["mosint"]