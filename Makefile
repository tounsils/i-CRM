# Quality gates. `make gates` is what has to pass before anything is pushed.
#
# Two ways in: the plain targets use whatever Go is on PATH, and the docker-
# prefixed ones run the same commands inside a pinned container for a machine
# with no Go toolchain installed.

GO              ?= go
GO_IMAGE        ?= golang:1.26-alpine
LINT_IMAGE      ?= golangci/golangci-lint:latest-alpine
MODULE_CACHE    ?= icrm-gomod
BINARY          ?= bin/icrm

# Mounts the working tree and a named volume for the module cache, so repeated
# runs do not re-download dependencies.
DOCKER_GO = docker run --rm \
	-v "$(CURDIR)":/src \
	-v $(MODULE_CACHE):/go/pkg/mod \
	-w /src $(GO_IMAGE)

.PHONY: help gates build lint format format-check test test-race run clean \
        docker-gates docker-build docker-up docker-down tidy

help:
	@echo "gates         build, format-check, lint and test"
	@echo "build         compile a static binary to $(BINARY)"
	@echo "format        rewrite files with gofmt"
	@echo "lint          golangci-lint, in Docker"
	@echo "test          go test with race detection off, for speed"
	@echo "test-race     go test -race"
	@echo "run           run the server locally on :8080"
	@echo "docker-gates  run the gates inside $(GO_IMAGE)"
	@echo "docker-build  build the deployable image"
	@echo "docker-up     bring up docker-compose"

gates:
	sh scripts/gates.sh
	$(MAKE) lint

build:
	CGO_ENABLED=0 $(GO) build -trimpath -ldflags="-s -w" -o $(BINARY) ./cmd/icrm

format:
	$(GO) run mvdan.cc/gofumpt@latest -l -w . 2>/dev/null || gofmt -l -w .

format-check:
	@unformatted=$$(gofmt -l ./cmd ./internal); \
	if [ -n "$$unformatted" ]; then \
		echo "gofmt would change:"; echo "$$unformatted"; exit 1; \
	fi
	@echo "format ok"

lint:
	docker run --rm -v "$(CURDIR)":/src -v $(MODULE_CACHE):/go/pkg/mod \
		-w /src $(LINT_IMAGE) golangci-lint run ./...

test:
	$(GO) vet ./...
	$(GO) test ./... -count=1

test-race:
	$(GO) test ./... -race -count=1

run:
	$(GO) run ./cmd/icrm

tidy:
	$(GO) mod tidy

clean:
	rm -rf bin data

# Same gates, for a machine without a Go toolchain.
docker-gates:
	$(DOCKER_GO) sh scripts/gates.sh

docker-build:
	docker build -t icrm:local .

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down
