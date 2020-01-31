GO_PATH=`go env GOPATH`

deps:
	go mod tidy

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/proxy main.go

help:
	@make2help $(MAKEFILE_LIST)
