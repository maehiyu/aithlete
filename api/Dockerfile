FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache tzdata
RUN go build -o api-server main.go

FROM alpine:3.18
WORKDIR /app
RUN apk add --no-cache tzdata
COPY --from=builder /app/api-server .
EXPOSE 8000
ENV PORT=8000
CMD ["./api-server"]
