
BINARY_DIRECTORY = "bin/movierental"

build:
	go build -o ${BINARY_DIRECTORY} ./cmd

run:
	go run cmd/main.go

test:
	go test -v ./...
