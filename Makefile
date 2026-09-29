include .env
export

export PROJECT_ROOT = $(CURDIR)

# ─── Docker ───────────────────────────────────────────────
up: ## Поднять postgres + migrate
	docker compose up -d

down: ## Остановить контейнеры (том с БД сохраняется)
	docker compose down

down-v: ## Остановить + удалить том с БД (сброс данных)
	docker compose down -v

# ─── Миграции ─────────────────────────────────────────────
migrate: ## Накатить миграции
	docker compose run --rm migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		up

migrate-down: ## Откатить последнюю миграцию
	docker compose run --rm migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		down 1

migrate-create: ## Создать миграцию: make migrate-create NAME=add_users
	docker compose run --rm migrate \
		create -ext sql -dir /migrations -seq $(NAME)

# ─── Приложение ───────────────────────────────────────────
run: ## Запустить приложение локально
	go run cmd/app/main.go
