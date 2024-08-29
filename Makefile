DB_URL ?= postgres://postgres@localhost:5432/golang_project_development?sslmode=disable
MIGRATIONS_PATH = db/migrate

migrate_up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate_down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate_version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version
