FROM alpine:3.14 AS builder

RUN mkdir -p /usr/src \
    && cd /usr/src \
    && wget -q https://github.com/php/php-src/archive/refs/tags/php-8.0.8.tar.gz \
    && tar -xzf php-8.0.8.tar.gz \
    && mv php-src-php-8.0.8 php

WORKDIR /usr/src/php

RUN apk add alpine-sdk autoconf automake libtool \
    && apk add bison re2c \
    && ./buildconf --force

RUN ./configure --disable-all --disable-cgi --disable-phpdbg --disable-debug CFLAGS="-O3 -march=native"

RUN make \
    && make install

RUN apk add upx \
    && upx -9 /usr/local/bin/php

FROM scratch

COPY --from=builder /usr/local/bin/php /usr/local/bin/php
COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1

COPY ./router.php /app/router.php

ENTRYPOINT ["/usr/local/bin/php"]

CMD ["-S", "0.0.0.0:8000", "/app/router.php"]

EXPOSE 8000
