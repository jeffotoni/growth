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
