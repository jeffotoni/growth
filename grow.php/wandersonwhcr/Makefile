PHP_VERSION="8.0.8"
PECL_SWOOLE_VERSION="4.7.0"
GIT_SHA_SHORT=$(shell git rev-parse --short HEAD)

.PHONY: all
all: build

.PHONY: base
base:
	docker build . \
		--file Base.Dockerfile \
		--build-arg PHP_VERSION=${PHP_VERSION} \
		--build-arg PECL_SWOOLE_VERSION=${PECL_SWOOLE_VERSION} \
		--tag wandersonwhcr/php:${PHP_VERSION}-cli

.PHONY: dev
dev: base
	docker run --rm --interactive --tty \
		--publish 8080:8080 \
		--volume `pwd`:/app:ro \
		wandersonwhcr/php:${PHP_VERSION}-cli

.PHONY: build
build: base
	docker build . \
		--file Dockerfile \
		--build-arg DOCKER_IMAGE_TAG=wandersonwhcr/php:${PHP_VERSION}-cli \
		--tag wandersonwhcr/growth:${GIT_SHA_SHORT}
	docker tag wandersonwhcr/growth:${GIT_SHA_SHORT} wandersonwhcr/growth:latest

.PHONY: stage
stage: build
	docker run --rm --interactive --tty \
		--publish 8080:8080 \
		wandersonwhcr/growth:${GIT_SHA_SHORT}
