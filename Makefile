
MIGRATE_CC = migrate  
DATABASE_URL = "postgres://test:test@localhost:4444/test?sslmode=disable"
MIGRATIONS_PATH= ./pkg/database/migrations
MIGRATIONS_SOURCE= file:./pkg/database/migrations

.PHONY: create-migration migrate-source migrate-down migrate-fixed

create-migration:
	$(MIGRATE_CC) create -ext sql -dir $(MIGRATIONS_PATH) -seq migration

migrate-up:
	$(MIGRATE_CC) -source $(MIGRATIONS_SOURCE) -database $(DATABASE_URL) -verbose up

migrate-down:
	@read -p "enter version migration: " version;\
	$(MIGRATE_CC) -source $(MIGRATIONS_SOURCE) -database $(DATABASE_URL) -verbose down $$version

migrate-fixed:
	@read -p "enter version migration: " version;\
	$(MIGRATE_CC) -source $(MIGRATIONS_SOURCE) -database $(DATABASE_URL) -verbose force $$version