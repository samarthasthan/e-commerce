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
up-prod:
	@echo "Running in Production mode..."
	@docker compose -f ./build/compose/compose.prod.yml up
	@echo "Production mode completed."

# Run in Development mode
up-dev:
	@echo "Running in Development mode..."
	@docker compose -f ./build/compose/compose.dev.yml up
	@echo "Development mode completed."

# Down in Production mode
down-prod:
	@echo "Running in Production mode..."
	@docker compose -f ./build/compose/compose.prod.yml up
	@echo "Production mode completed."

# Down in Development mode
down-dev:
	@echo "Running in Development mode..."
	@docker compose -f ./build/compose/compose.dev.yml up
	@echo "Development mode completed."

# Make migrations
migrate-up:
	@echo "Making migrations..."
	@migrate -path ./internal/authentication/database/mysql/migrations -database "mysql://root:password@tcp(localhost:3306)/authentication" -verbose up
	@echo "Migrations completed."

# Delete migrations
migrate-down:
	@echo "Deleting migrations..."
	@migrate -path ./internal/authentication/database/mysql/migrations -database "mysql://root:password@tcp(localhost:3306)/authentication" -verbose down	@echo "Migrations deleted."
