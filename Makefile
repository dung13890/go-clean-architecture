%.sql:
	$(eval TS := $(shell date '+%Y%m%d%H%M%S'))
	touch db/migrations/${TS}_$(*F).up.sql
	touch db/migrations/${TS}_$(*F).down.sql

lint:
	golangci-lint run --config=.golangci.yaml

test: ### run test
	go test -v -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -func cover.out

.PHONY: go-gen lint test
# put the files with interfaces you'd like to mock in prerequisites
# wildcards are allowed

go-gen:
	@echo "Generating mocks..."
	rm -rf internal/domain/mock/*_mock.go
	go generate -x ./internal/...
	@echo "Generating pkg mocks..."
	rm -rf pkg/*/mock/*_mock.go
	go generate -x ./pkg/...

dev:
	air -c cmd/app/.air.toml
