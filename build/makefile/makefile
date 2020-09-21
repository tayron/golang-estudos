PWD := $(shell pwd)

compile:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go

compress:
	@upx main

build: compile compress

.PHONY: compile