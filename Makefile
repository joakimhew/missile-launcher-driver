EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))


BINARY=hatch
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Build=${BUILD}"

default: build

all: clean build_all install ## Runs clean, build_all and install

test: ## Runs all available tests
	go test ./...

fmt: ## Formats everything
	go fmt ./...

build_driver: ## Builds driver
	go build ${LDFLAGS} $(CMD_DIR)

# Remove only what we've created
clean: ## Cleans the binary
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

.PHONY: check clean all

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


# Example: make generate-proto proto=driver output=./cmd/iotcore lang=node
.PHONY: generate-proto
generate-proto: 
	docker run -v `pwd`:/defs namely/protoc-all:1.28_1 -i ./api/protobuf -f "${proto}/${proto}.proto" -o ${output}/api -l ${lang}