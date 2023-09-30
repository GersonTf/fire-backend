# Default listen address
listenaddr ?= :8080

# Build the project
build:
	go build -o fire-backend

# Run the project
run:
	go run main.go --listenaddr $(listenaddr)
