mongo_uri = mongodb://localhost:27017

build:
	@go build -o bin/habit-tracker-api .

run: build
	@MONGO_URI=${mongo_uri} ./bin/habit-tracker-api

unit:
	@MONGO_URI=${mongo_uri} go test -v ./... -short

integration: 
	@MONGO_URI=${mongo_uri} go test -v ./... -run Integration

test:
	@MONGO_URI=${mongo_uri} go test -v ./...

local-mongo lm:
	@docker-compose -f local-mongo-docker-compose.yml up -d --build --remove-orphans

local-mongo-stop lm-stop:
	@docker-compose -f local-mongo-docker-compose.yml down
