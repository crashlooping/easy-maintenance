# easy-maintenance

https://github.com/crashlooping/easy-maintenance

A small application to provide a http endpoint for maintenance tasks.

## Build and run

```bash
# Windows
go build -o easy-maintenance-app.exe

# Linux
CGO_ENABLED=0 go build -o easy-maintenance-app .

# Docker
docker build -t easy-maintenance-app .
docker run --rm -p 8080:8080 easy-maintenance-app
```

## ghcr.io

```bash
docker pull ghcr.io/crashlooping/easy-maintenance/easy-maintenance:latest
docker run --rm -p 8080:8080 ghcr.io/crashlooping/easy-maintenance/easy-maintenance:latest
```

## Go

```bash
go get -u
go mod tidy -go=1.21
```
