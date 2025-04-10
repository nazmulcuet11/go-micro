FRONT_BINARY=frontApp
BROKER_BINARY=brokerApp
LOGGER_BINARY=loggerApp
AUTH_BINARY=authApp
MAIL_BINARY=mailApp
LISTENER_BINARY=listenerApp
FRONT_END_BINARY=frontEndApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker compose (if running), builds all projects and starts docker compose
up_build: build_broker build_logger build_auth build_mail build_listener
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## build all binaries
build_all: build_broker build_logger build_auth build_mail build_listener build_front_end

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/${BROKER_BINARY}-arm64 ./cmd/api
	cd ./broker-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/${BROKER_BINARY}-amd64 ./cmd/api
	@echo "Done!"

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/${LOGGER_BINARY}-arm64 ./cmd/api
	cd ./logger-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/${LOGGER_BINARY}-amd64 ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ./authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/${AUTH_BINARY}-arm64 ./cmd/api
	cd ./authentication-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/${AUTH_BINARY}-amd64 ./cmd/api
	@echo "Done!"

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo "Building mail binary..."
	cd ./mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/${MAIL_BINARY}-arm64 ./cmd/api
	cd ./mail-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/${MAIL_BINARY}-amd64 ./cmd/api
	@echo "Done!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ./listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/${LISTENER_BINARY}-arm64 .
	cd ./listener-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/${LISTENER_BINARY}-amd64 .
	@echo "Done!"

## build_front_end: builds the front end binary as a linux executable
build_front_end:
	@echo "Building front end binary..."
	cd ./front-end && env GOOS=linux CGO_ENABLED=0 go build -o ./build/${FRONT_END_BINARY}-arm64 ./cmd/web
	cd ./front-end && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/${FRONT_END_BINARY}-amd64 ./cmd/web
	@echo "Done!"

## build_front: builds the front end binary
build_front:
	@echo "Building front end binary..."
	cd ./front-end && env CGO_ENABLED=0 go build -o ./build/${FRONT_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ./front-end && ./build/${FRONT_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./build/${FRONT_BINARY}"
	@echo "Stopped front end!"