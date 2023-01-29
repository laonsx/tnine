all: auth game gate

auth:
	go build -o ./bin/auth ./cmd/auth

gate:
	go build -o ./bin/gate ./cmd/gate

game:
	go build -o ./bin/game ./cmd/game

run:
	./bin/auth rpc
	./bin/auth api
	./bin/game rpc
	./bin/game api
	./bin/gate rpc
	./bin/gate api