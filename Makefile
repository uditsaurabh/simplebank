createdb:
	docker exec -it postgres_container psql -U udit -d jobber_reviews -c "CREATE DATABASE jobber_reviews;"

dropdb:
	docker exec -it postgres_container psql -U udit -d jobber_reviews -c "DROP DATABASE jobber_reviews;"

migrateup:
	migrate -path db/migration -database "postgresql://udit:root@localhost:5432/jobber_reviews?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://udit:root@localhost:5432/jobber_reviews?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test