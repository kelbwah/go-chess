run:
	@./bin/go-chess
test:
	@go test -v ./...
build:
	@go build -o ./bin/go-chess ./cmd/main.go
