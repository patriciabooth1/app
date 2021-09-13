bin             := executable

.EXPORT_ALL_VARIABLES:
GO111MODULE = on
CGO_ENABLED = 0

.PHONY: all
all: build

.PHONY: clean
clean:
	go mod tidy
	go fmt ./...
	rm -f $(bin)

.PHONY: test-unit
test-unit:
	go test -v ./... -run Unit

.PHONY: build
build: clean
	go build -o ./$(bin)