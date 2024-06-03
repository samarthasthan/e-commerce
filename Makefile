#Makefile

up:
	@echo "Uping all services"
	@rm -rf frontend/admin-portal/node_modules frontend/admin-portal/build frontend/seller-portal/node_modules frontend/seller-portal/.next frontend/website/node_modules frontend/website/.next
	@docker-compose -f builds/package/compose.yaml up -d

down:
	@echo "Downing all services"
	@docker-compose -f builds/package/compose.yaml down --volumes

build:
	@echo "Building all services"
	@rm -rf frontend/admin-portal/node_modules frontend/admin-portal/build frontend/seller-portal/node_modules frontend/seller-portal/.next frontend/website/node_modules frontend/website/.next
	@docker-compose -f builds/package/compose.yaml build

grpc-gen-all: grpc-gen grpc-gen-js

grpc-clean-all: grpc-clean grpc-clean-js

grpc-gen:
	@echo "Generating Go code from Protocol Buffers..."
	@protoc --go_out=paths=source_relative:./pkg/proto_go --go-grpc_out=paths=source_relative:./pkg/proto_go --proto_path=./pkg/proto ./pkg/proto/*.proto
	@echo "Go code generation completed."

grpc-clean:
	@echo "Cleaning generated Go code..."
	@rm -f ./pkg/proto_go/*.go
	@echo "Clean completed."

grpc-gen-js:
	@echo "Generating JavaScript code from Protocol Buffers..."
	@protoc -I=proto proto/*.proto \
		--js_out=import_style=commonjs:proto_js \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:proto_js
	@echo "JavaScript code generation completed."

grpc-clean-js:
	@echo "Cleaning generated JavaScript code..."
	@rm -f proto_js/*.js
	@echo "Clean completed."
	
run-authentication:
	@go run cmd/authentication/main.go

run-broker:
	@go run cmd/broker/main.go
