## Запуск в докере

### Создать `.env.docker`

```bash
touch task_tracker_app/.env.docker
```

#### Пример переменных окружения
```
POSTGRES_USER=example_user
POSTGRES_PASSWORD=example_pass
POSTGRES_DB=golang_task_tracker

DB_NAME=golang_task_tracker
DB_HOST=db
DB_PORT=5432
DB_USER=example_user
DB_PASS=example_pass
DB_SSL_MODE=disable

JWT_EXPIRATION=600
YANDEX_SPELLER_URL=https://speller.yandex.net/services/spellservice.json/checkText

DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)
```

### Заупск контейнеров

```bash
docker-compose up --build
```

### Прогоняем миграции
```bash
make ENV=docker docker_migrate_up
```