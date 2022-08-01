FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN GOOS=linux go build -o api ./cmd/api

FROM alpine:latest

# Copy binary from builder
COPY --from=builder /app/api /usr/bin/

EXPOSE 8081

ENTRYPOINT ["/usr/bin/api"]
