FROM golang:1.24-bookworm AS builder

RUN useradd -u 1001 nonroot

WORKDIR /app

COPY go.mod .

ENV GIN_MODE=release

RUN go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o http-handler \
    ./cmd/http-handler

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /app/config/development.toml /config/config.toml

COPY --from=builder /app/http-handler /http-handler

USER nonroot

CMD ["./http-handler"]
