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
