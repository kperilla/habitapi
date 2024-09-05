build:
	@go build -o bin/habit-tracker-api .

run: build
	@./bin/habit-tracker-api

test:
	@go test -v ./...
