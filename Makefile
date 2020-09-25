.PHONY: build
build: generate
	go build -mod=vendor -o server ./cmd/server/main.go

.PHONY: generate
generate:
	go generate ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test: test/server test/frontend

.PHONY: test/server
test/server:
	go test -mod=vendor -race -v -cover $$(go list ./... )

.PHONY: test/frontend
test/frontend:
	cd web && yarn test

.PHONY: lint
lint: lint/server lint/frontend

.PHONY: lint/server
lint/server:
	gofmt -d internal/ cmd/
	revive \
		-config ./revive.toml \
		-exclude ./vendor/... \
		-exclude ./internal/database/internal/... \
		-formatter stylish \
		./...

.PHONY: lint/frontend
lint/frontend:
	cd web && yarn lint
