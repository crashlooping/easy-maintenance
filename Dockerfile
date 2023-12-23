FROM alpine:3
RUN apk update --no-cache && \
    apk upgrade --no-cache
WORKDIR /app
COPY html html
COPY easy-maintenance-app ./

EXPOSE 8080

ENTRYPOINT ["/app/easy-maintenance-app"]
