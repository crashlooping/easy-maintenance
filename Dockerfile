
FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go build -o easy-maintenance-app

FROM alpine:latest
RUN apk update --no-cache && \
    apk upgrade --no-cache
WORKDIR /app
COPY html html
COPY --from=builder /build/easy-maintenance-app ./

ENTRYPOINT ["/app/easy-maintenance-app"]
