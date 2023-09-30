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
	docker run -p 8080:8080 fire-backend