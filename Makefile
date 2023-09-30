IMAGE_NAME=fire-backend

# Default listen address
listenaddr ?= :8080

# Build the project
build:
	go build -o fire-backend

# Run the tests
tests:
	go test ./...

# Run the project
run:
	go run main.go --listenaddr $(listenaddr)

# Build docker image
docker-build:
	docker build -t fire-backend .

# Run docker container
docker-run:
	docker run -e FIRE_JWT_SECRET='$(FIRE_JWT_SECRET)' -e MONGO_URI='$(MONGO_URI)' -p 8080:8080 $(IMAGE_NAME)