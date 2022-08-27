%.sql:
	$(eval TS := $(shell date '+%Y%m%d%H%M%S'))
	touch db/migrations/${TS}_$(*F).up.sql
	touch db/migrations/${TS}_$(*F).down.sql

lint:
	golangci-lint run --config=.golangci.yaml

test: ### run test
	go test -v -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -func cover.out

.PHONY: build-mock
# put the files with interfaces you'd like to mock in prerequisites
# wildcards are allowed

build-mock:
	@echo "Generating mocks..."
	mockgen -source=internal/domain/user.go -destination=internal/domain/mock/user_mock.go
	mockgen -source=internal/domain/role.go -destination=internal/domain/mock/role_mock.go
	mockgen -source=internal/domain/auth.go -destination=internal/domain/mock/auth_mock.go
