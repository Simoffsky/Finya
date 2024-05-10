.PHONY: run, build

run: build
	./bin/bot

build:
	go build -o ./bin/bot cmd/bot/main.go 
test:
	go test ./... -race