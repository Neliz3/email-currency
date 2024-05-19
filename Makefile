build:
	docker-compose build ecurrency

run:
	docker-compose up ecurrency


migrate_up:
	migrate -path=internal/database/migrations -database 'postgres://postgres:$(PASSWORD)@localhost:5436/postgres?sslmode=disable' -verbose up

migrate_down:
	migrate -path=internal/database/migrations -database 'postgres://postgres:$(PASSWORD)@localhost:5436/postgres?sslmode=disable' -verbose down

migrate_version:
	migrate -path=internal/database/migrations -database 'postgres://postgres:$(PASSWORD)@localhost:5436/postgres?sslmode=disable' version

migrate_force:
	migrate -path=internal/database/migrations -database 'postgres://postgres:$(PASSWORD)@localhost:5436/postgres?sslmode=disable' force $(VERSION)



docker_db:
	docker exec -it $(ID) /bin/bash

