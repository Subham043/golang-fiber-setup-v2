.PHONY: clean critic security lint test build run \
        docker.up docker.down \
        start stop start.server stop.server \
        start.app.windows start.app.linux \
        stop.app.windows stop.app.linux \
        swag

APP_NAME = golang_fiber_setup
ROOT_DIR := $(CURDIR)
BUILD_DIR = $(ROOT_DIR)/tmp
# MIGRATIONS_FOLDER = $(ROOT_DIR)/platform/migrations
# DATABASE_URL = mysql://root:go_pwd@go_mysql/go_db?sslmode=disable

ifeq ($(OS),Windows_NT)
	START_APP := start.app.windows
	STOP_APP := stop.app.windows
else
	START_APP := start.app.linux
	STOP_APP := stop.app.linux
endif

ifeq ($(OS),Windows_NT)
clean:
	@if exist $(BUILD_DIR) rmdir /s /q $(BUILD_DIR)
else
clean:
	rm -rf $(BUILD_DIR)
endif

ifeq ($(OS),Windows_NT)
APP_BIN := $(APP_NAME).exe
else
APP_BIN := $(APP_NAME)
endif

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: 
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_BIN) main.go

run: build
	$(BUILD_DIR)/$(APP_BIN)

docker.up: 
	docker compose -f ./docker/docker-compose.yml up -d

docker.down: 
	docker compose -f ./docker/docker-compose.yml down

start.app.windows: 
	air -c .air.windows.conf

start.app.linux: 
	air -c .air.linux.conf

start.server:
	$(MAKE) $(START_APP)

stop.app.windows:
	air -c .air.windows.conf -s

stop.app.linux:
	air -c .air.linux.conf -s

stop.server:
	$(MAKE) $(STOP_APP)

start:
	$(MAKE) docker.up
	$(MAKE) start.server

stop:
	$(MAKE) docker.down
	$(MAKE) stop.server

swag:
	swag init

check.air:
	@command -v air >/dev/null 2>&1 || { echo "❌ air not installed"; exit 1; }

# Tidy
tidy:
	go mod tidy

# Generate Ent boilerplate
generate:
	go generate ./cmd