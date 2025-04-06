# Fire-Backend Project

Welcome to the Fire-Backend project.

## Prerequisites

- Go 1.24.0 or later
- Docker (for running tests with test containers)
- Make (optional, but recommended for using the Makefile)

## Running the Fire Backend Project

This project is based on Go 1.24.0. You can build and run it using the Makefile commands:

### Building the Project

```bash
make build
```

### Running the Project

```bash
make run
```

The default listen address is: `:8080`

### Dockerized Version

The project is also dockerized and can be built and run using Docker.

#### Building the Docker Image

```bash
make docker-build
```

#### Running the Docker Image

```bash
make docker-run
```

## Development Tools

### Linting

This project uses [`golangci-lint`](https://golangci-lint.run/usage/quick-start/) as its linter.

#### Installation

To install the latest version of `golangci-lint`:

```bash
make install-lint
```

#### Usage

To run the linter:

```bash
make lint
```

Or directly:

```bash
golangci-lint run
```

### Testing

The project uses test containers for integration tests with MongoDB. Make sure Docker is running before executing tests.

#### Running Tests

To run all tests:

```bash
make test
```

This will run tests with verbose output and proper test container support.

#### Additional Test Options

You can also run tests directly with Go commands:

```bash
# Run all tests in all packages
go test -v ./...

# Run with test coverage
go test -cover ./...

# Run tests in random order
go test -v -shuffle=on ./...
```

### Running All Checks

To run both linting and tests:

```bash
make check
```

## Documentation

To access the project documentation:

1. Install the Go documentation tool:

   ```bash
   go install golang.org/x/tools/cmd/godoc
   ```

2. Run the documentation server:

   ```bash
   godoc -play -http ":6060"
   ```

3. Access the documentation at:
   - General Go documentation: `http://localhost:6060/pkg/`
   - Project documentation: `http://localhost:6060/pkg/github.com/GersonTf/fire-backend/`
