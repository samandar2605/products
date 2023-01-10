POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_DATABASE=products

-include .env
  
DB_URL=postgresql://postgres:12345@localhost:5432/products?sslmode=disable



swag:
	swag init -g api/api.go -o api/docs

run:
	go run "./cmd/main.go"

mock:
	mockgen -package mockdb -destination storage/mockdb/storage.go mocking/storage StorageI

migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up


migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY: run swag migrateup migratedown mock