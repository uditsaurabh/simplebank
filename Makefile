createdb:
	docker exec -it postgres_container psql -U udit -d jobber_reviews -c "CREATE DATABASE jobber_reviews;"

dropdb:
	docker exec -it postgres_container psql -U udit -d jobber_reviews -c "DROP DATABASE jobber_reviews;"

migrateup:
	migrate -path db/migration -database "postgresql://udit:root@localhost:5432/jobber_reviews?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://udit:root@localhost:5432/jobber_reviews?sslmode=disable" -verbose up 1


migratedown:
	migrate -path db/migration -database "postgresql://udit:root@localhost:5432/jobber_reviews?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://udit:root@localhost:5432/jobber_reviews?sslmode=disable" -verbose down 1


sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination orm/mock/store.go  github.com/uditsaurabh/go-simple-bank/orm Store	

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1