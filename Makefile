.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o reverse-proxy-go main.go
.PHONY:build

run: build
	./reverse-proxy-go config.json
.PHONY:run