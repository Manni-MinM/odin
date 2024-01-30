FROM golang:alpine AS builder

WORKDIR /app
COPY ./ ./

WORKDIR /app/cmd/odin-healthcheck
RUN go build -o odin-healthcheck

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/odin-healthcheck/odin-healthcheck .

ENTRYPOINT ["./odin-healthcheck"]
