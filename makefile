all: run-importer run-api

run-importer:
	DUMP_FILE=data_dump.csv \
	DB_USER=root \
	DB_PASS=password \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_SCHEMA=geolocation \
	GRPC_SERVER_HOST=localhost \
	GRPC_SERVER_PORT=8080 \
	go run ./cmd/importer/

run-api:
	echo "TBD"

generate-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./handler/grpc/schema/schema.proto

integration-tests:
	docker compose up -d
	DB_USER=root \
	DB_PASS=password \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_SCHEMA=geolocation \
	go test -v -tags=integration ./...