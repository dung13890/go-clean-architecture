%.sql:
	$(eval TS := $(shell date '+%Y%m%d%H%M'))
	touch db/migrations/${TS}_$(*F).up.sql
	touch db/migrations/${TS}_$(*F).down.sql

lint:
	golangci-lint run --config=.golangci.yaml

test: ### run test
	go test -v -cover -race ./...
