FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ws-server main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/ws-server .
EXPOSE 9000
CMD ["./ws-server"]
