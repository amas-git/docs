ARG GO_VERSION=1.14
ARG VERSION
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src/app
COPY . /src/app
RUN echo "version="$VERSION
RUN CGO_ENABLED=0 GOOS=linux VERSION=${VERSION} go build -ldflags "-X main.version=$VERSION" cmd/echosvc.go


FROM alpine:latest
LABEL maintainer=zhoujb.cn@gmail.com
USER nobody
WORKDIR /app
EXPOSE ${PORT:-8888} 8080
COPY --from=builder --chown=nobody /src/app/cert    ./cert
COPY --from=builder --chown=nobody /src/app/echosvc .
ENTRYPOINT [ "/app/echosvc" ]