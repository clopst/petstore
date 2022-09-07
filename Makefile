postgres:
		docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
		docker exec -it postgres14 createdb --username=postgres --owner=postgres petstores

dropdb:
		docker exec -it postgres14 dropdb -U postgres petstores

migrate:
		migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/petstores?sslmode=disable" -verbose up

.PHONY: postgres createdb dropdb migrate