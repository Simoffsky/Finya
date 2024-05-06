.PHONY: run

run:
	go run ./cmd/bot/main.go

test:
	go test ./... -race