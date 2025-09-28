FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 9000
ENV PORT=9000

CMD ["./main"]