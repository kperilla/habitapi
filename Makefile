ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@go build -o bin/habit-tracker-api .

run: build
	./bin/habit-tracker-api

unit:
	go test -v ./... -short

integration itest int:
	go test -v ./... -run Integration

test:
	go test -v ./...

local-mongo lm:
	@docker-compose -f local-mongo-docker-compose.yml up -d --build --remove-orphans

local-mongo-stop lm-stop:
	@docker-compose -f local-mongo-docker-compose.yml down
