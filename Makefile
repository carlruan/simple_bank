postgres:
	docker run --name postgres12 --network bank_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=19981024 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:yD4Ri1WL2aTB7tLFu0yA@simplebanl.comurmkyanli.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:19981024@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:19981024@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:19981024@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/carlruan/simple_bank/db/sqlc Store

.PHONT: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1