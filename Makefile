up:
	rm -rf frontend/admin-portal/node_modules frontend/admin-portal/build frontend/seller-portal/node_modules frontend/seller-portal/.next frontend/website/node_modules frontend/website/.next
	docker compose -f builds/package/compose.yaml up -d

down:
	docker compose -f builds/package/compose.yaml down --volumes