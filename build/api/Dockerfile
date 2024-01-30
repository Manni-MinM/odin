FROM golang:alpine AS builder

WORKDIR /app
COPY ./ ./

WORKDIR /app/cmd/odin-api
RUN go build -o odin-api

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/odin-api/odin-api .

EXPOSE 8000
ENTRYPOINT ["./odin-api"]
