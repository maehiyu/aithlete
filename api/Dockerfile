FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api-server main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/api-server .
EXPOSE 8000
ENV PORT=8000
CMD ["./api-server"]
