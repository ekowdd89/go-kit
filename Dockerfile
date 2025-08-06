ARG GOLANG=golang:1.23



FROM ${GOLANG} AS base

ENTRYPOINT [ "go", "run" ]
CMD [ "./cmd/author" ]



FROM base AS debugger

ENTRYPOINT [ "go", "run", "-mod=mod", "github.com/go-delve/delve/cmd/dlv@latest"]
CMD [ "debug", "./cmd/author", "--headless", "--listen=:2345", "--accept-multiclient", "--continue", "--build-flags='-buildvcs=false'" , "--api-version=2"]



FROM base AS builder

WORKDIR /app
COPY ./ ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/tmp/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o author ./cmd/author



FROM alpine AS alpine

RUN apk add --no-cache ca-certificates



FROM scratch

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/author /usr/local/bin/author

ENTRYPOINT ["/usr/local/bin/author"]