ENV_FILE_DOCKER_PATH=task_tracker_app/.env.docker
ENV_FILE_DEV_PATH=task_tracker_app/.env.dev
ENV ?= local

ifeq ($(ENV), docker)
    include $(ENV_FILE_DOCKER_PATH)
    export $(shell sed 's/=.*//' $(ENV_FILE_DOCKER_PATH))
else
    include $(ENV_FILE_DEV_PATH)
    export $(shell sed 's/=.*//' $(ENV_FILE_DEV_PATH))
endif

DB_URL ?= postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)
MIGRATIONS_PATH = task_tracker_app/db/migrate

migrate_up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate_down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate_version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

docker_migrate_up:
	docker compose -f docker-compose.yml --profile tools run --rm migrate up