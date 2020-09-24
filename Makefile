.PHONY: build
build: generate
	go build -mod=vendor -o server ./cmd/server/main.go

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test -mod=vendor -race -v -cover $$(go list ./... )

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	gofmt -d internal/ cmd/
	revive \
		-config ./revive.toml \
		-exclude ./vendor/... \
		-exclude ./internal/database/internal/... \
		-formatter stylish \
		./...
