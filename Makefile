postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=4d45b61f-5df9-4568-a912-d89d2e7c13a2 -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:4d45b61f-5df9-4568-a912-d89d2e7c13a2@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup-last:
	migrate -path db/migration -database "postgresql://root:4d45b61f-5df9-4568-a912-d89d2e7c13a2@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:4d45b61f-5df9-4568-a912-d89d2e7c13a2@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown-last:
	migrate -path db/migration -database "postgresql://root:4d45b61f-5df9-4568-a912-d89d2e7c13a2@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	moq -pkg mockdb -out db/mock/store.go db/sqlc Querier:MockStore

.PHONY: postgres createdb dropdb migrateup migratedown migrateup-last migratedown-last sqlc server