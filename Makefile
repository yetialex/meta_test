BUILD_FOLDER := dist

build: ## Build binary and docker images
	docker build -t build.meta --force-rm -f build.Dockerfile .
	mkdir -p $(CURDIR)/$(BUILD_FOLDER)
	docker run -v $(CURDIR)/$(BUILD_FOLDER):/opt/mount --rm --entrypoint cp build.meta /src/app/dist/meta-server /opt/mount/meta-server
	docker build -t meta --force-rm -f Dockerfile .

run:
	docker-compose up -d

#######################################################################
# Codegen

generate:
	find ./gen/swagger -mindepth 1 ! -name "configure_meta.go" -name "meta-server" -delete || true && \
		${GOPATH}/bin/swagger027 generate server -f ./swagger-test/swagger.json -A meta --target=./gen/swagger

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
META_SERVER_NAME = "meta-server"
META_SERVER_IMAGE_BASE=${META_SERVER_NAME}/web
META_SERVER_IMAGE=${META_SERVER_IMAGE_BASE}:$(APP_VERSION)
META_SERVER_IMAGE_LATEST=${META_SERVER_IMAGE_BASE}:latest

docker-build-meta-server: docker-pull-golang
	docker build -t ${META_SERVER_IMAGE} -f ./docker/Dockerfile .

docker-run-meta-server: docker-build-meta-server docker-remove-meta-server
	docker run -d \
		--net ${NETWORK} \
		--name ${META_SERVER_NAME} \
		-e LOGGING_LEVEL=debug \
		${META_SERVER_IMAGE}

docker-remove-meta-server:
	docker rm -f ${META_SERVER_NAME} || true

docker-run-all: \
	docker-remove-all \
	docker-create-network \
	docker-run-postgres \
	docker-run-meta-server

docker-remove-all: \
	docker-remove-meta-server \
	docker-remove-postgres \
	docker-remove-network
