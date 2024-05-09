build:
	@go build -o ./bin/ ./cmd/...

run:
	@go run ./cmd/server/main.go

build-run: build
	@./bin/server

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./pkg/proto/exchange.proto

