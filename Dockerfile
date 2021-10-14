# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# RUN apk add --no-cache add build-base git gcc bash

WORKDIR /workspace/go

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./ 
# COPY . . 

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o /workspace/go/golang_redis

EXPOSE 8080

ENTRYPOINT ["/golang_redis"]
CMD ["serve"]