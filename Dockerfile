FROM alpine:latest
RUN apk update --no-cache && \
    apk upgrade --no-cache
WORKDIR /app
COPY html html
COPY easy-maintenance-app ./

ENTRYPOINT ["/app/easy-maintenance-app"]
