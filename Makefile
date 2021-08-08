
.PHONY: build
build:
	go build -o dist/main cmd/main.go

run: build
	dist/main

.PHONY: clean
clean:
	rm -rf dist/*

.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate types -o internal/api/ports/openapi_types.gen.go -package ports api/openapi/links.yml
	oapi-codegen -generate chi-server -o internal/api/ports/openapi_api.gen.go -package ports api/openapi/links.yml
