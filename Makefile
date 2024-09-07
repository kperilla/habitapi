build:
	@go build -o bin/habit-tracker-api .

run: build
	@MONGO_URI=mongodb://localhost:27017 ./bin/habit-tracker-api

test:
	@go test -v ./...

local-mongo lm:
	@docker-compose -f local-mongo-docker-compose.yml up -d --build --remove-orphans

local-mongo-stop lm-stop:
	@docker-compose -f local-mongo-docker-compose.yml down
