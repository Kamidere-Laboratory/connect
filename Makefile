all: fmt vet mod lint gen

# Run tests
test: fmt vet
	go test ./...

# Run go fmt against code
fmt:
	go fmt ./...

# Run go fmt against code
mod:
	go mod tidy && go mod verify

# Run go vet against code
vet:
	go vet ./...

# Run linters
lint:
	golangci-lint run

# Serve the docs website locally and auto on changes
dev-docs:
	cd .web && yarn install && yarn dev
