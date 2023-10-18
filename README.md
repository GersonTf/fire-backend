# Fire-Backend Project

Welcome to the Fire-Backend project.

## Running the Fire Backend Project

This project is based on Go 1.21.X You can build and run it as follows:

### Building the Project

```
go build -o fire-backend
```

### Running the Project

Execute the following command:

```
go run main.go --listenaddr $(listenaddr)
```

The default listen address is: `8080`

### Dockerized Version

The project is also dockerized and can be built and run using Docker.

#### Building the Docker Image

Set the image name:

```
IMAGE_NAME=fire-backend
```

Then, build the image:

```
docker build -t $(IMAGE_NAME) .
```

#### Running the Docker Image

```
docker run -p 8080:8080 $(IMAGE_NAME)
```

## Linting with `golangci-lint`

This project uses [`golangci-lint`](https://golangci-lint.run/usage/quick-start/) as its linter.

### Installation

To install `golangci-lint` locally, you can use the following command:

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
```

After installation, you can verify the version with:

```bash
golangci-lint --version
```

### Usage

To run the lint tests, use the following command:

```bash
golangci-lint run
```

This is equivalent to:

```bash
golangci-lint run ./...
```

For those who prefer Docker, you can use the Docker version as follows:

```bash
docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v
```

to run the tests:
go test ./...

add -v (verbose) for more detailed test logs
-shuffle=on for executing tests in a random order

## Running Tests

To execute the tests, use the following command:

```
go test ./...
```

### Additional Options

1. **Verbose Output:** To view more detailed test logs, use the `-v` flag.

   ```
   go test ./... -v
   ```

2. **Randomized Test Execution:** If you wish to execute tests in a random order, include the `-shuffle=on` flag.

   ```
   go test ./... -shuffle=on
   ```

3. **Specific Test Case:** To run a particular test case, utilize the `run` flag followed by the test name. The `run` flag accepts regex patterns, enabling you to target specific or groups of tests.

   ```
   go test ./... -run=<test_pattern>
   ```

Replace `<test_pattern>` with your desired regex pattern or test name.
