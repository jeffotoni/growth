FROM alpine:3.14 AS builder

ARG PHP_VERSION

RUN mkdir -p /usr/src \
    && cd /usr/src \
    && wget -q https://github.com/php/php-src/archive/refs/tags/php-${PHP_VERSION}.tar.gz \
    && tar -xzf php-${PHP_VERSION}.tar.gz \
    && mv php-src-php-${PHP_VERSION} php

RUN cd /usr/src/php/ext \
    && wget -q https://pecl.php.net/get/apcu-5.1.20.tgz \
    && tar -xzf apcu-5.1.20.tgz \
    && mv apcu-5.1.20 apcu

WORKDIR /usr/src/php

RUN apk add alpine-sdk autoconf automake libtool \
    && apk add bison re2c \
    && ./buildconf --force

RUN ./configure --disable-all --disable-cgi --disable-phpdbg --disable-debug --enable-apcu CFLAGS="-O3 -march=native"

RUN make \
    && make install

RUN apk add upx \
    && upx -9 /usr/local/bin/php

FROM scratch

COPY --from=builder /usr/local/bin/php /usr/local/bin/php
COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1

ENTRYPOINT ["/usr/local/bin/php"]

CMD ["-S", "0.0.0.0:8080", "/app/router.php"]

WORKDIR "/app"

EXPOSE 8080
