# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-alpine AS build

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
RUN go build -o /template

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /workspace/go

COPY --from=build /template /template

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["template"]
CMD ["serve"]