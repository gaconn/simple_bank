postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank
dropdb:
	docker exec -it postgres12 dropdb simplebank
migrateup:
	migrate -path db/migration -database postgres://root:secret@localhost:5432/simplebank?sslmode=disable -verbose up
migratedown:
	migrate -path db/migration -database postgres://root:secret@localhost:5432/simplebank?sslmode=disable -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
start:
	go run main.go
mock: 
	mockgen -package mockdb  -destination db/mock/store.go  github.com/quan12xz/simple_bank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test start
