FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ai-coach main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/ai-coach .
EXPOSE 9100
ENV PORT=9100
CMD ["./ai-coach"]
