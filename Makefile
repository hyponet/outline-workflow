.PHONY: all build clean run check cover lint docker help
BIN_DIR=bin
BASE_PATH=$(shell pwd)
all: check build
build:
	docker run --rm -v $(BASE_PATH):/go/src/github.com/hyponet/outline-workflow \
	-v $(BASE_PATH)/bin:/bin/outline-workflow \
	-w /go/src/github.com/hyponet/outline-workflow \
	golang:1.16 sh ./hack/multibuild.sh ./cmd /bin/outline-workflow
clean:
	@go clean
test:
	@go test ./...
check:
	@go fmt ./...
	@go vet ./...
lint:
	golangci-lint run --enable-all
help:
	@echo "make build - build multi arch binary"
	@echo "make clean - clean workspace"
	@echo "make test  - run all testcase"
	@echo "make check - go format and vet"
	@echo "make lint  - golint"