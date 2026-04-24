TARGET := $(word 2, $(MAKECMDGOALS))

.PHONY: create-migration 
create-migration: # create your migrations
	@echo "Generating a new migration!"
	migrate create -ext sql -dir ./internal/infraestructure/db/sqlc/migrations $(TARGET)
	

.PHONY: migration 
migration: #up or down your migrations.
	migrate -source file://./internal/infraestructure/db/sqlc/migrations -database "mysql://root:root@tcp(localhost:3306)/database_v1_monolith" -verbose $(TARGET)