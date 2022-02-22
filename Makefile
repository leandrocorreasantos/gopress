IMAGE=app
DOCKER_CMD=docker exec -it ${IMAGE}

_create-env:
	[ -f .env ] || cp .env.sample .env

build:
	docker-compose up -d

init:
	${DOCKER_CMD} go mod init ${IMAGE}
	${DOCKER_CMD} go mod tidy

setup: _create-env
	${DOCKER_CMD} go mod tidy
	${DOCKER_CMD} go get

start:
	${DOCKER_CMD} go run main.go

compile:
	${DOCKER_CMD} go mod tidy
	${DOCKER_CMD} go get
	${DOCKER_CMD} go build ${IMAGE}

log:
	docker logs ${IMAGE}

test:
	${DOCKER_CMD} go test
