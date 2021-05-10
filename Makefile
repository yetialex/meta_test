THIS_FILE := $(lastword $(MAKEFILE_LIST))

#######################################################################
# Codegen

generate:
	find ./gen/swagger -mindepth 1 ! -name "configure_meta.go" -delete || true && \
		${GOPATH}/bin/swagger027 generate server -f ./swagger-test/swagger.json -A meta --target=./gen/swagger

build-app:
	go build -o bin/meta-app ./gen/swagger/cmd/meta-server

build-app-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/meta-app ./gen/swagger/cmd/meta-server

run-app:
		LOGGING_LEVEL=debug \
        DB_URL="postgres://postgres:secret@localhost/metaservice?sslmode=disable&pool_max_conns=10" \
        PORT=${APP_PORT} \
		./bin/meta-app

#######################################################################
# Docker

NETWORK="meta"

docker-remove-network:
	docker network remove ${NETWORK} || true

docker-create-network: docker-remove-network
	docker network create ${NETWORK}

#######################################################################
# Pull postgres image
POSTGRES_IMAGE=postgres:latest

docker-pull-postgres:
	docker pull ${POSTGRES_IMAGE}

docker-remove-postgres:
	docker rm -f postgres || true

docker-run-postgres: docker-pull-postgres docker-remove-postgres
	docker run -d \
        --net=${NETWORK} \
        --name=postgres \
        -p 5432:5432 \
        -e POSTGRES_PASSWORD=secret \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_DB=meta \
        -v /var/lib/docker/volumes/meta_test_pg_project/_data:/var/lib/postgresql/data \
        ${POSTGRES_IMAGE}

# Pull base golang image
docker-pull-golang:
	docker pull golang:latest

#######################################################################
# Meta server
APP_VERSION = 0.0.1
APP_PORT = "8092"
APP_HOST = "0.0.0.0"
META_SERVER_NAME = "meta-server"
META_SERVER_IMAGE_BASE=${META_SERVER_NAME}/web
META_SERVER_IMAGE=${META_SERVER_IMAGE_BASE}:$(APP_VERSION)
META_SERVER_IMAGE_LATEST=${META_SERVER_IMAGE_BASE}:latest

docker-build-app: docker-pull-golang
	docker build -t ${META_SERVER_IMAGE} -f ./docker/Dockerfile .

docker-run-app: docker-build-app docker-remove-app
	docker run -d \
		--net ${NETWORK} \
		--name ${META_SERVER_NAME} \
		-p ${APP_PORT}:${APP_PORT} \
		-e PORT=${APP_PORT} \
		-e HOST=${APP_HOST} \
		-e LOGGING_LEVEL=debug \
		-e DB_URL="postgres://postgres:secret@postgres/metaservice?sslmode=disable&pool_max_conns=10" \
		${META_SERVER_IMAGE}

docker-remove-app:
	docker rm -f ${META_SERVER_NAME} || true

docker-run-all: \
	docker-remove-all \
	docker-create-network \
	docker-run-postgres \
	docker-run-app

docker-remove-all: \
	docker-remove-app \
	docker-remove-postgres \
	docker-remove-network

#######################################################################
# Tests

tests-dep:
	cd ./tests && \
        python3 -m venv venv && \
		. venv/bin/activate && \
		pip install -r requirements.txt --upgrade && \
    	deactivate

tests-run:
	cd ./tests && \
		. venv/bin/activate && \
		pytest -W ignore -vv && \
		deactivate

tests-sleep: tests-dep docker-run-all
	sleep 5

tests-run-app: tests-sleep
	@$(MAKE) -f $(THIS_FILE) docker-run-app

test-run-with-docker: tests-run-app
	cd ./tests && \
		. venv/bin/activate && \
		pytest -W ignore -vv && \
		deactivate
