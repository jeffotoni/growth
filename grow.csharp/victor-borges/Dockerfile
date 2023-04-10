FROM alpine:3.17 AS base
WORKDIR /app
EXPOSE 8080

ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1 \
    DOTNET_RUNNING_IN_CONTAINER=true \
    ASPNETCORE_URLS=http://+:8080

RUN apk add --no-cache libstdc++

FROM mcr.microsoft.com/dotnet/sdk:8.0-preview-alpine3.17 AS build
WORKDIR /src

COPY "Growth.csproj" .
RUN dotnet restore "Growth.csproj"
COPY . .
RUN dotnet build "Growth.csproj" -c Release

FROM build AS publish
RUN dotnet publish "Growth.csproj" -c Release -o /app -r linux-musl-x64 \
    /p:DebugType=None /p:DebugSymbols=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
ENTRYPOINT ["./Growth"]
