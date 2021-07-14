# create binary
FROM golang:1.16.0 AS build-executable
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -a" -o exec

# compress the binary using upx
FROM alpine:latest as compressed
RUN apk add --no-cache upx
COPY --from=build-executable /src /src
WORKDIR /src
RUN upx exec

# Get the binary already compressed
FROM scratch
COPY --from=compressed /src /
ENTRYPOINT [ "/exec" ]