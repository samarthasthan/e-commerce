#Makefile

up:
	@echo "Uping all services"
	@rm -rf frontend/admin-portal/node_modules frontend/admin-portal/build frontend/seller-portal/node_modules frontend/seller-portal/.next frontend/website/node_modules frontend/website/.next
	@docker compose -f builds/package/compose.yaml up -d

down:
	@echo "Downing all services"
	@docker compose -f builds/package/compose.yaml down --volumes

build:
	@echo "Building all services"
	@rm -rf frontend/admin-portal/node_modules frontend/admin-portal/build frontend/seller-portal/node_modules frontend/seller-portal/.next frontend/website/node_modules frontend/website/.next
	@docker compose -f builds/package/compose.yaml build

gen:
	@echo "Generating Go code from Protocol Buffers..."
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
	@echo "Go code generation completed."

clean:
	@echo "Cleaning generated Go code..."
	@rm -f proto/*.go
	@echo "Clean completed."
