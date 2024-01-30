FROM golang:alpine AS builder

WORKDIR /app
COPY ./ ./

WORKDIR /app/cmd/odin-cron
RUN go build -o odin-cron

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/odin-cron/odin-cron .

ENTRYPOINT ["./odin-cron"]
