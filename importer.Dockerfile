FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN GOOS=linux go build -o importer ./cmd/importer

FROM alpine:latest

# Copy binary from builder
COPY --from=builder /app/importer /usr/bin/
# Copy of the datadump file
COPY --from=builder /app/data_dump.csv /

EXPOSE 8080

ENTRYPOINT ["/usr/bin/importer"]
