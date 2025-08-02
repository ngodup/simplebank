postgres:
	docker run --name pg-root-container -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:16

createdb:
	docker exec -it pg-root-container createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -e PGPASSWORD=mysecretpassword -it pg-root-container dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

# test:
# 	go test -v -cover ./...
test:
	go test -v -count=1 -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ngodup/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
