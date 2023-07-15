build:
	go build -o ./out/app cmd/api/main.go

run:
	go run cmd/api/main.go

test: test-unit test-integration

test-unit:
	go test -v -tags=unit ./internal/...

test-integration:
	go test -v ./test/api/...

codegen:
	oapi-codegen -generate types,spec -package models definition/api.yaml > internal/api/models/models.gen.go
	oapi-codegen -generate gin -package services definition/api.yaml > internal/api/services/service.gen.go
