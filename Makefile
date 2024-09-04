build:
	@go build -o bin/habit-tracker-api ./cmd

run: build
	@./bin/habit-tracker-api

test:
	@go test -v ./...
