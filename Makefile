.PHONY: build-local build templ notify-templ-proxy run

build-local:
	@go build -o ./tmp/web ./cmd/web

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./tmp/web ./cmd/web

templ:
	@go tool templ generate --watch --proxy=http://localhost:8080 --proxyport=8082 --open-browser=false --proxybind="0.0.0.0"

notify-templ-proxy:
	@go tool templ generate --notify-proxy --proxyport=8082

run:
	@make templ & sleep 1
	@air
