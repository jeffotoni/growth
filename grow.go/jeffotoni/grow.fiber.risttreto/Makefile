# Makefile
.EXPORT_ALL_VARIABLES:	

#GO111MODULE=on
#GOPROXY=direct
#GOSUMDB=off
GOPRIVATE=github.com/jeffotoni/growth

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o fiberisttreto main.go
	upx fiberisttreto
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/growth/grow.go/jeffotoni/grow.fiber.risttreto
	go mod tidy
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o fiberisttreto main.go
	upx fiberisttreto
	@echo "buid update completo..."
	@echo "fim"

docker:
	@docker build --no-cache -f Dockerfile -t growth/gofiberisttreto .
