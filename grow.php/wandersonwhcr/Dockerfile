ARG DOCKER_IMAGE_TAG

FROM ${DOCKER_IMAGE_TAG} AS base

RUN apk add upx \
    && upx -9 /usr/local/bin/php

FROM scratch

COPY --from=base /usr/local/bin/php /usr/local/bin/php
COPY --from=base /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1

COPY ./router.php /app/router.php
COPY ./src /app/src

WORKDIR "/app"

ENTRYPOINT ["/usr/local/bin/php"]

CMD ["-S", "0.0.0.0:8080", "router.php"]

EXPOSE 8080
