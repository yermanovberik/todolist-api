.PHONY: build run docker-build docker-run docker-compose-up

build:
	go build -o main ./cmd

run: build
	./main

docker-build:
	docker build -t todo-list-api .

docker-run:
	docker run -p 8080:8080 todo-list-api

docker compose up:
	docker compose up --build