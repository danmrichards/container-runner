GOARCH=amd64

.PHONY: build
build: build-server build-runner

.PHONY: build-server
build-server:
	GOOS=linux go build -ldflags="-s -w" -o bin/server-linux-${GOARCH} ./cmd/server/main.go

.PHONY: build-runner
build-runner:
	GOOS=linux go build -ldflags="-s -w" -o bin/runner-linux-${GOARCH} ./cmd/runner/main.go

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

.PHONY: clean
clean:
	rm -rf rootfs.tar

rootfs.tar: clean
	./rootfs.sh
