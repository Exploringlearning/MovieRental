
BINARY_DIRECTORY = "bin/movierental"
HOME_PATH = $(shell echo $$HOME)
export DOCKER_HOST=unix://${HOME_PATH}/.colima/default/docker.sock
export TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE=/var/run/docker.sock


build:
	go build -o ${BINARY_DIRECTORY} ./cmd

run-without-build:
	go run cmd/main.go

run: build
	./${BINARY_DIRECTORY}

test: unit-test integration-test

unit-test:
	go test ./internal/...

integration-test:
	go test -v ./test/integration_test/...

dep:
	go mod download

clean:
	go clean
	rm ${BINARY_DIRECTORY}
