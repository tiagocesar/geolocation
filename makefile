

run-importer:
	DUMP_FILE=data_dump.csv \
	DB_USER=root \
	DB_PASS=password \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_SCHEMA=geolocation \
	go run ./cmd/importer/main.go

generate-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./handler/grpc/schema/schema.proto
