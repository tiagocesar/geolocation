FROM golang:alpine as builder

WORKDIR /app

COPY . .

# For the datadump file
COPY data_dump.csv .

RUN GOOS=linux go build -o importer ./cmd/importer

FROM alpine:latest

# Copy binary from builder
COPY --from=builder /app/importer /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/importer"]
