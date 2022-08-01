run:
	COMPOSE_PROFILES=backend docker compose up -d --build

stop:
	COMPOSE_PROFILES=backend docker compose down --rmi all -v

run-importer:
	DUMP_FILE=data_dump.csv \
	DB_USER=root \
	DB_PASS=password \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_SCHEMA=geolocation \
	GRPC_SERVER_PORT=8080 \
	go run ./cmd/importer/

run-api:
	HTTP_SERVER_PORT=8081 \
	GRPC_SERVER_HOST=localhost \
	GRPC_SERVER_PORT=8080 \
	go run ./cmd/api/

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