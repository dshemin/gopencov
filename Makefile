.PHONY: build
build:
	go build -mod=vendor -o server ./cmd/server/main.go

.PHONY: test
test:
	go test -mod=vendor -race -v -cover $$(go list ./... )

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	gofmt -d `find . -type f -name '*.go' -not -path "./vendor/*"`
	revive \
		-config ./revive.toml \
		-exclude ./vendor/... \
		-formatter stylish \
		./...
