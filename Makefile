run:
	@export $$(grep -v '^#' .env | xargs) && go run ./cmd/api

tidy:
	go mod tidy

build:
	go build -o bin/api ./cmd/api

docker-build:
	docker build -t thistoric-api:local .

docker-run:
	docker run --rm -p 8080:8080 --env-file .env thistoric-api:local
