run-game:
	@./bin/go-chess --game
run-server:
	@./bin/go-chess --server
test:
	@go test -v ./...
build:
	@go build -o ./bin/go-chess ./cmd/main.go
