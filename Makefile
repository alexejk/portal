# We don't need make's built-in rules.
MAKEFLAGS += --no-builtin-rules

GO_FLAGS= CGO_ENABLED=0
GO_LDFLAGS= -ldflags=""
GO_BUILD_CMD=$(GO_FLAGS) go build $(GO_LDFLAGS)

BINARY_NAME=portal
BUILD_DIR=build

.PHONY: docker

all: clean generate-all lint test build-all

lint:
	@echo "Linting code..."
	@go vet ./...
test:
	@echo "Running tests..."
	@go test ./...

code-gen:
	@echo "Generating code..."
	@go generate ./...

generate-all: code-gen

pre-build:
	@mkdir -p $(BUILD_DIR)

build-linux: pre-build
	@echo "Building Linux binary..."
	GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64
build-osx: pre-build
	@echo "Building OSX binary..."
	GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64
build build-all: build-linux # build-osx

clean:
	@echo "Cleaning..."
	@rm -Rf $(BUILD_DIR)

docker:
# Build a new image (delete old one)
	docker build --force-rm --build-arg GOPROXY -t $(BINARY_NAME) .

build-in-docker: docker
# Force-stop any containers with this name
	docker rm -f $(BINARY_NAME) || true
# Create a new container with newly built image (but don't run it)
	docker create --name $(BINARY_NAME) $(BINARY_NAME)
# Copy over the binary to disk (from container)
	docker cp '$(BINARY_NAME):/opt/' $(BUILD_DIR)
# House-keeping: removing container
	docker rm -f $(BINARY_NAME)