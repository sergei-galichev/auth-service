include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIG_DIR=$(MIG_DIR)
LOCAL_MIG_DSN="host=$(M_HOST) port=$(M_PORT) dbname=$(M_DB_NAME) user=$(M_USER) password=$(M_PASS) sslmode=$(M_SSL)"

AUTH_BIN:=authApp

# INSTALL AND GET DEPENDENCIES
install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/pressly/goose/v3

# MIGRATIONS
migration-create:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIG_DIR} create migration_file sql

migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIG_DIR} postgres ${LOCAL_MIG_DSN} status -v

migration-up-all:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIG_DIR} postgres ${LOCAL_MIG_DSN} up -v

migration-up-by-one:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIG_DIR} postgres ${LOCAL_MIG_DSN} up-by-one -v

migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIG_DIR} postgres ${LOCAL_MIG_DSN} down -v

# GO TOOLS RUNNING
format:
	go fmt ./...

# GRPC API: generates from proto files

generate:
	make generate-auth-api

generate-auth-api:
	@mkdir -p pkg/grpc/v1/auth
	@echo "Generating stubs..."
	@protoc --proto_path api/v1 \
	--go_out=pkg/grpc/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/grpc/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/v1/auth/auth.proto
	@echo "Done! Generated."

# DOCKER
## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Done! Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_auth_service
	@echo "Stopping Docker images (if running)..."
	docker compose down
	@echo "Done! Docker compose stopped."
	@echo "Building (when required) and starting Docker images..."
	docker compose up --build -d
	@echo "Done! Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping Docker compose..."
	docker compose down
	@echo "Done! Docker compose stopped."

## build_auth_service: build the auth service binary as a linux executable
build_auth_service:
	@echo "Building auth service binary..."
	env GOOS=linux CGO_ENABLED=0 GOMAXPROCS=4 go build -o $(LOCAL_BIN)/$(AUTH_BIN) ./cmd/app
	@echo "Done! Auth service binary built."
