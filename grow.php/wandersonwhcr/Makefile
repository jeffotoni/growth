PHP_VERSION="8.0.8"
PECL_APCU_VERSION="5.1.20"
GIT_SHA_SHORT=$(shell git rev-parse --short HEAD)

.PHONY: all
all: build

.PHONY: build
build: base

.PHONY: base
base:
	docker build . \
		--file Base.Dockerfile \
		--build-arg PHP_VERSION=${PHP_VERSION} \
		--build-arg PECL_APCU_VERSION=${PECL_APCU_VERSION} \
		--tag wandersonwhcr/php:${PHP_VERSION}-cli

.PHONY: dev
dev: base
	docker run --rm --interactive --tty \
		--publish 8080:8080 \
		--volume `pwd`:/app:ro \
		wandersonwhcr/php:${PHP_VERSION}-cli

.PHONY: artifact
artifact: base
	docker build . \
		--file Artifact.Dockerfile \
		--build-arg PHP_VERSION=${PHP_VERSION} \
		--tag wandersonwhcr/growth:${GIT_SHA_SHORT}
