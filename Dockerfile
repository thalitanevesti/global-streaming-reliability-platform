# syntax=docker/dockerfile:1.7
FROM golang:1.27.1-bookworm AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go test ./... && \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w -X main.version=0.1.0" -o /out/streaming-api ./cmd/api

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build --chown=65532:65532 /out/streaming-api /streaming-api
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/streaming-api"]

