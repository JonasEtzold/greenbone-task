docs:
	swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o ./out/app cmd/api/main.go

run:
	go run cmd/api/main.go

build-docker: build
	docker build . -t <%= serviceName %>

run-docker: build-docker
	docker run -p 3000:3000 gostarter

test: test-unit test-integration

test-unit:
	go test -v -tags=unit ./internal/...

test-integration:
	go test -v ./test/api/...

codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	oapi-codegen -generate types -package api api/api.yaml > internal/api/models.gen.go
    oapi-codegen -generate server -package api api/api.yaml > internal/api/server.gen.go
