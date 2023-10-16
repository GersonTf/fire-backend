# Fire-Backend Project

Welcome to the Fire-Backend project.

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
