
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
	oapi-codegen -generate types -o pkg/links/ports/openapi_types.gen.go -package links api/openapi/links.yml
	oapi-codegen -generate chi-server -o pkg/links/ports/openapi_api.gen.go -package links api/openapi/links.yml

.PHONY: unit_test
unit_test:
	go test -p=8 ./...
