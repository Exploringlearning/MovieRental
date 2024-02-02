
BINARY_DIRECTORY = "bin/movierental"

build:
	go build -o ${BINARY_DIRECTORY} ./cmd

run-without-build:
	go run cmd/main.go

run: build
	./${BINARY_DIRECTORY}

# test:
#    export DOCKER_HOST="unix://${HOME}/.colima/default/docker.sock"
#    export TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE="/var/run/docker.sock"
# 	go test ./...

dep:
	go mod download

clean:
	go clean
	rm ${BINARY_DIRECTORY}
