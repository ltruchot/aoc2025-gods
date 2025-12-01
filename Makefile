.PHONY: serve build test lint generate clean install dev install-browsers

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/air-verse/air@latest
	go mod tidy

install-browsers:
	go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps chromium

generate:
	templ generate

build: generate
	go build -o bin/server ./cmd/server

serve: generate
	go run ./cmd/server

dev:
	templ generate -watch -proxy="http://localhost:8080" -cmd="go run ./cmd/server" -open-browser=false -watch-pattern="(.+\\.go$$)|(.+\\.templ$$)|(.+\\.css$$)"

lint:
	go fmt ./...
	templ fmt .

test:
	go test ./tests/e2e/... -v

clean:
	rm -rf bin/
	rm -rf test-results/
	find . -name "*_templ.go" -delete
