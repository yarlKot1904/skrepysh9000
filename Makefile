.PHONY: init
init:
	go install golang.org/x/tools/cmd/goimports@latest
	go mod tidy

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -C skrepysh-agent -o bin/skrepysh-agent -trimpath cmd/main.go

.PHONY: format
format:
	gofmt -s -w .
	goimports -w -local skrepysh-agent .
