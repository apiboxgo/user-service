BINARY_NAME=main.go
ENV_FILES=.env .env.dev
DB_URL_DEV=postgres://postgres:postgres@localhost:5432/apigobox?sslmode=disable
DB_URL_TEST=postgres://postgres:postgres@localhost:5432/apigobox-test?sslmode=disable
MIGRATIONS_DIR=./db/migrations
TESTS_DIR=./api/user
SERVER_PORT=8081
default: run

build:
	go build -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)
.PHONY: build

run:
	@echo "▶️  Running $(BINARY_NAME)..."
	APP_ENV=dev go run $(BINARY_NAME)
.PHONY: run

runs:
	@echo "▶️  Running with swagger $(BINARY_NAME)..."
	swag init && APP_ENV=dev go run $(BINARY_NAME)
.PHONY: runs

stop:
	@echo "Остановка сервера"
	@kill -SIGINT $(shell lsof -t -i:$(SERVER_PORT))
.PHONY: stop
fmt:
	go fmt ./...
.PHONY: fmt

tidy:
	go mod tidy
.PHONY: tidy

test:
	APP_ENV=test go test $(TESTS_DIR)
.PHONY: test

testv:
	APP_ENV=test go test $(TESTS_DIR) -v
.PHONY: testv

clean:
	rm -f $(BINARY_NAME)
.PHONY: clean

migration-create:
	goose create -dir $(MIGRATIONS_DIR) $(NAME) sql
.PHONY: migration-create

migration-reset:
	goose  -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_DEV)" reset
.PHONY: migration-reset

migration-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_DEV)" up
.PHONY: migration-up

migration-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_DEV)" down
.PHONY: migration-down

migration-up-test:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_TEST)" up
.PHONY: migration-up-test

migration-down-test:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_TEST)" down
.PHONY: migration-down-testt

migration-reset-test:
	goose  -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_TEST)" reset
.PHONY: migration-reset-test