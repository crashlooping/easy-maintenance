FROM alpine:3
RUN apk update --no-cache && \
    apk upgrade --no-cache
WORKDIR /app
COPY html html
COPY static static
COPY easy-maintenance-app ./

EXPOSE 8080

ENTRYPOINT ["./easy-maintenance-app"]