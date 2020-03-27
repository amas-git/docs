FROM golang:1.14-alpine AS builder
WORKDIR /src/app
COPY . /src/app
RUN go build cmd/...


FROM golang:1.14-alpine
WORKDIR /app
COPY --from=builder /src/app/echosvc /app