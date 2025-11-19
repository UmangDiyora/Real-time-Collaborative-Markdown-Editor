.PHONY: help build run test clean migrate-up migrate-down docker-build docker-up docker-down

# Variables
APP_NAME=markdown-collab
MAIN_PATH=./cmd/server
MIGRATE_PATH=./cmd/migrate
BUILD_DIR=./bin
DOCKER_COMPOSE=docker-compose

# Default target
help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make migrate-up     - Run database migrations"
	@echo "  make migrate-down   - Rollback database migrations"
	@echo "  make migrate-create - Create new migration (use NAME=migration_name)"
	@echo "  make docker-build   - Build Docker images"
	@echo "  make docker-up      - Start Docker containers"
	@echo "  make docker-down    - Stop Docker containers"
	@echo "  make lint           - Run linters"
	@echo "  make fmt            - Format code"

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/server $(MAIN_PATH)
	@go build -o $(BUILD_DIR)/migrate $(MIGRATE_PATH)
	@echo "Build complete!"

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_PATH)/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete!"

# Run migrations up
migrate-up:
	@echo "Running migrations..."
	@go run $(MIGRATE_PATH)/main.go up

# Run migrations down
migrate-down:
	@echo "Rolling back migrations..."
	@go run $(MIGRATE_PATH)/main.go down

# Create new migration
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	@go run $(MIGRATE_PATH)/main.go create $(NAME)

# Docker commands
docker-build:
	@echo "Building Docker images..."
	@$(DOCKER_COMPOSE) build

docker-up:
	@echo "Starting Docker containers..."
	@$(DOCKER_COMPOSE) up -d

docker-down:
	@echo "Stopping Docker containers..."
	@$(DOCKER_COMPOSE) down

docker-logs:
	@$(DOCKER_COMPOSE) logs -f

# Development helpers
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete!"

lint:
	@echo "Running linters..."
	@golangci-lint run ./...

# Install development dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed!"

# Watch for changes and rebuild (requires air)
dev:
	@command -v air > /dev/null 2>&1 || { echo "air not installed. Install with: go install github.com/cosmtrek/air@latest"; exit 1; }
	@air
