.PHONY: build run test docker-build docker-run

build:
	go build -o bin/app cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

docker-build:
	docker build -t app .

docker-run:
	docker compose up
