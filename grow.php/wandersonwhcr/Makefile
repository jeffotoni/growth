PHP_VERSION="8.0.8"

.PHONY: all
all: build

.PHONY: build
build: base

.PHONY: base
base:
	docker build . \
		--build-arg PHP_VERSION=${PHP_VERSION} \
		--tag wandersonwhcr/php:${PHP_VERSION}-cli

.PHONY: dev
dev: base
	docker run --rm --interactive --tty \
		--publish 8080:8080 \
		--volume `pwd`:/app:ro \
		wandersonwhcr/php:${PHP_VERSION}-cli
