ARG GO_VERSION=1.14
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src/app
COPY . /src/app
RUN CGO_ENABLED=0 GOOS=linux go build cmd/echosvc.go


FROM alpine:latest
LABEL maintainer=zhoujb.cn@gmail.com
USER nobody
WORKDIR /app
EXPOSE ${PORT:-8888}
COPY --from=builder --chown=nobody /src/app/cert    /app/cert
COPY --from=builder --chown=nobody /src/app/echosvc /app
ENTRYPOINT [ "/app/echosvc" ]