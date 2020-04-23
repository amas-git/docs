ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /app
COPY .  /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "" cmd/hello.go



FROM alpine:latest
LABEL maintainer=zhoujb.cn@gmail.com
USER nobody
WORKDIR /app
EXPOSE 6666
COPY --from=builder --chown=nobody /app/hello .