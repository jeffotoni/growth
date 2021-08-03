FROM alpine:3.14

ARG PHP_VERSION
ARG PECL_APCU_VERSION

RUN mkdir -p /usr/src \
    && cd /usr/src \
    && wget -q https://github.com/php/php-src/archive/refs/tags/php-${PHP_VERSION}.tar.gz \
    && tar -xzf php-${PHP_VERSION}.tar.gz \
    && mv php-src-php-${PHP_VERSION} php

RUN cd /usr/src/php/ext \
    && wget -q https://pecl.php.net/get/apcu-${PECL_APCU_VERSION}.tgz \
    && tar -xzf apcu-${PECL_APCU_VERSION}.tgz \
    && mv apcu-${PECL_APCU_VERSION} apcu

WORKDIR /usr/src/php

RUN apk add alpine-sdk autoconf automake libtool \
    && apk add bison re2c \
    && ./buildconf --force

RUN ./configure --disable-all \
        --disable-cgi \
        --disable-phpdbg --disable-debug \
        --enable-apcu \
        CFLAGS="-O3 -march=native" \
    && sed -i 's/-export-dynamic/-all-static/g' Makefile

RUN make \
    && make install

WORKDIR "/app"

ENTRYPOINT ["/usr/local/bin/php"]

CMD ["-S", "0.0.0.0:8080", "router.php"]

EXPOSE 8080
