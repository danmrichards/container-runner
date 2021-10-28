GOARCH=amd64

.PHONY: build
build: build-server build-manager

.PHONY: build-server
build-server:
	GOOS=linux go build -ldflags="-s -w" -o bin/server-linux-${GOARCH} ./cmd/server/main.go

.PHONY: build-manager
build-manager:
	GOOS=linux go build -ldflags="-s -w" -o bin/manager-linux-${GOARCH} ./cmd/manager/main.go

.PHONY: lint
lint:
	golangci-lint run ./cmd/...

.PHONY: test
test:
	go test -v -race -count=1 ./...

.PHONY: deps
deps:
	go mod verify && \
	go mod tidy
