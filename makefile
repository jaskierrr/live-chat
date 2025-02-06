.PHONY: run

include .env

run:
		go run cmd/app/main.go

build:
		docker-compose up --build -d

test:
		go test -v -count=1 ./...

models:
		reform-db -db-driver=postgres -db-source="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -models=./internal/models init

env:
		echo

venv:
  source .venv/bin/activate
