#Makefile

# Generate Go code from Protocol Buffers
grpc-gen:
	@echo "Generating Go code from Protocol Buffers..."
	@protoc --go_out=paths=source_relative:./pkg/proto_go --go-grpc_out=paths=source_relative:./pkg/proto_go --proto_path=./pkg/proto ./pkg/proto/*.proto
	@echo "Go code generation completed."

# Clean generated Go code
grpc-clean:
	@echo "Cleaning generated Go code..."
	@rm -f ./pkg/proto_go/*.go
	@echo "Clean completed."

# Run in Production mode
prod-up:
	@echo "Running in Production mode..."
	@make migrate-up
	@docker compose -f ./build/compose/compose.prod.yaml up -d
	@echo "Production mode completed."


# Down in Production mode
prod-down:
	@echo "Running in Production mode..."
	@docker compose -f ./build/compose/compose.prod.yaml down
	@echo "Production mode completed."

# Run in Development mode
dev-up:
	@echo "Running in Development mode..."
	@docker compose -f ./build/compose/compose.dev.yaml up -d --no-deps --build
	@make migrate-up
	@echo "Development mode completed."

# Reload in Development mode
dev-reload:
	@echo "Reloading in Development mode..."
	@docker compose -f ./build/compose/compose.dev.yaml pull
	@docker compose -f ./build/compose/compose.dev.yaml up -d --build
	@echo "Reload completed."
	
# Down in Development mode
dev-down:
	@echo "Running in Development mode..."
	@make migrate-down
	@docker compose -f ./build/compose/compose.dev.yaml down --volumes
	@echo "Development mode completed."

# Make migrations
migrate-up:
	@echo "Making migrations..."
	@migrate -path ./internal/authentication/database/mysql/migrations -database "mysql://root:password@tcp(localhost:3306)/authentication" -verbose up
	@echo "Migrations completed."

# Delete migrations
migrate-down:
	@echo "Deleting migrations..."
	@migrate -path ./internal/authentication/database/mysql/migrations -database "mysql://root:password@tcp(localhost:3306)/authentication" -verbose down

# SQLC generate
sqlc-gen:
	@echo "Generating SQLC..."
	@sqlc generate -f ./internal/authentication/database/mysql/sqlc/sqlc.yaml
	@echo "SQLC generation completed."

# Unit tests
unit-test:
	@echo "Running unit tests..."
	@go test -v ./...
	@echo "Unit tests completed."