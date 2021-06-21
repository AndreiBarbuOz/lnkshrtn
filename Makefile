
.PHONY: build
build:
	go build -o dist/main main.go

run: build
	dist/main

.PHONY: clean
clean:
	rm -rf dist/*

