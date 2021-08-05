ARG DOCKER_IMAGE_TAG

FROM ${DOCKER_IMAGE_TAG} AS base

RUN apk add upx \
    && upx -9 /usr/local/bin/php

FROM scratch

COPY --from=base /usr/local/bin/php /usr/local/bin/php

COPY ./server.php /app/server.php
COPY ./src /app/src

WORKDIR "/app"

ENTRYPOINT ["/usr/local/bin/php"]

CMD ["server.php"]

EXPOSE 8080
