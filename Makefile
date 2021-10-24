
.PHONY: build-server
build-server:
	go build -o dist/lnk-server cmd/server/main.go

run: build-server
	dist/lnk-server

.PHONY: build-cli
build-cli:
	go build -o dist/lnkctl cmd/cli/main.go

.PHONY: clean
clean:
	rm -rf dist/*

.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate types -o pkg/links/ports/openapi_types.gen.go -package links api/openapi/links.yml
	oapi-codegen -generate chi-server -o pkg/links/ports/openapi_api.gen.go -package links api/openapi/links.yml

	oapi-codegen -generate types -o pkg/apiclient/openapi_types.gen.go -package apiclient api/openapi/links.yml
	oapi-codegen -generate client -o pkg/apiclient/openapi_api.gen.go -package apiclient api/openapi/links.yml

.PHONY: unit-test
unit-test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./pkg/... ./cmd/cli/...

.PHONY: build_instrumented
build_instrumented:
	go test -c -o dist/lnk-server-dbg -covermode=atomic -coverpkg=all ./cmd/server
