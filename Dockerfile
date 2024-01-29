FROM golang:alpine AS builder

WORKDIR /app
COPY ./ ./

RUN go build -o odin

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/odin .

EXPOSE 8000
ENTRYPOINT ["./odin"]
