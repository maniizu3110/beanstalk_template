postgres:
	docker run --name postgres12 --network template_default -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
create-container:
	docker build -t template:latest .
	docker run --name template -p 8080:8080 --network template_default -e DB_SOURCE="postgresql://root:secret@postgres12:5432/template?sslmode=disable" template:latest
# docker-reset作成する
createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root template

dropdb:
	docker exec -it postgres12 dropdb template

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/template?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/template?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/template?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/template?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

run:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  server/api/sqlc Store
.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock run mock createmigrate