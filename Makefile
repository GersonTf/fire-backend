IMAGE_NAME=fire-backend

# Default listen address
listenaddr ?= :8080

# Build the project
build:
	go build -o fire-backend

# Run the tests
tests:
	go test ./...

# run the lint
lint:
	golangci-lint run

# Run the project
run:
	go run main.go --listenaddr $(listenaddr)

# Build docker image
docker-build:
	docker build -t $(IMAGE_NAME) .

# Run docker container
docker-run:
	docker run \
	-e FIRE_JWT_SECRET='$(FIRE_JWT_SECRET)' \
	-e DB_NAME='$(DB_NAME)' \
	-e MONGO_URI='$(MONGO_URI)' \
	-p 8080:8080 $(IMAGE_NAME)
