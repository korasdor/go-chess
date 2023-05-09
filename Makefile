.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans app

debug: build
	docker-compose up --remove-orphans debug

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out | grep "total"

lint:
	golangci-lint run

migrate:
	migrate -path ./schema -database 'postgres://postgres:root@0.0.0.0:5432/chess_db?sslmode=disable' up

swag:
	swag init -g internal/app/app.go