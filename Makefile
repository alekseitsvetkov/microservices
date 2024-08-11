.PHONY: compose-up migrate-create migrate-up

compose-up:
	docker-compose -f $(FILE) up -d

migrate-create:
	@set "GOOSE_MIGRATION_DIR=.\migrations" && go run github.com/pressly/goose/v3/cmd/goose postgres "user=microservices password=microservices dbname=microservices sslmode=disable" create $(name) sql

migrate-up:
	@set "GOOSE_MIGRATION_DIR=.\migrations" && go run github.com/pressly/goose/v3/cmd/goose postgres "user=microservices password=microservices dbname=microservices sslmode=disable" up

migrate-down:
	@set "GOOSE_MIGRATION_DIR=.\migrations" && go run github.com/pressly/goose/v3/cmd/goose postgres "user=microservices password=microservices dbname=microservices sslmode=disable" down

migrate-up-to:
	@set "GOOSE_MIGRATION_DIR=.\migrations" && go run github.com/pressly/goose/v3/cmd/goose postgres "user=microservices password=microservices dbname=microservices sslmode=disable" up $(name)

migrate-status:
	@set "GOOSE_MIGRATION_DIR=.\migrations" && go run github.com/pressly/goose/v3/cmd/goose postgres "user=microservices password=microservices dbname=microservices sslmode=disable" status

migrate-reset:
	@set "GOOSE_MIGRATION_DIR=.\migrations" && go run github.com/pressly/goose/v3/cmd/goose postgres "user=microservices password=microservices dbname=microservices sslmode=disable" reset
