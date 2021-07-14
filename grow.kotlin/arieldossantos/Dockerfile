FROM gradle:6.8.3-jre11 as builder

RUN mkdir /app

WORKDIR /app

COPY . .

RUN chown -R gradle /app

RUN ./gradlew installDist || return 0


FROM alpine:latest as packager

RUN apk --no-cache add openjdk11-jdk openjdk11-jmods

ENV JAVA_MINIMAL="/opt/java-minimal"

# build minimal JRE
RUN /usr/lib/jvm/java-11-openjdk/bin/jlink \
    --verbose \
    --add-modules \
        java.base,java.sql,java.naming,java.desktop,java.management,java.security.jgss,java.instrument \
    --compress 2 --strip-debug --no-header-files --no-man-pages \
    --release-info="add:IMPLEMENTOR=radistao:IMPLEMENTOR_VERSION=radistao_JRE" \
    --output "$JAVA_MINIMAL"



FROM alpine:latest

ENV JAVA_HOME=/opt/java-minimal

ENV PATH="$PATH:$JAVA_HOME/bin"

COPY --from=packager "$JAVA_HOME" "$JAVA_HOME"

RUN mkdir /app

WORKDIR /app

COPY --from=builder "/app/build/install/com.growth/" "/app"

EXPOSE 8080

CMD [ "sh", "/app/bin/com.growth" ]
