DB_NAME=ecom
ROOT_USER=root
ROOT_PASSWORD=securepass

build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

startdb:
	@docker run --name ecom-db -p 3307:3306 \
	-e MYSQL_ROOT_PASSWORD=${ROOT_PASSWORD} \
	-d mysql; sleep 20;
	@docker exec -it -e MYSQL_PWD=${ROOT_PASSWORD} ecom-db \
	mysql -u ${ROOT_USER} -e 'CREATE DATABASE IF NOT EXISTS ${DB_NAME};'

stopdb:
	@docker rm ecom-db -f

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

.PHONY: build run test startdb stopdb migration migrate-up migrate-down