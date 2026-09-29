## API
- POST   /users          — создать пользователя
- GET    /users/{uuid}   — получить по uuid
- DELETE /users/{uuid}   — удалить по uuid

## Запуск
make up        # поднять Postgres
make migrate   # накатить миграции
make run       # запустить сервер
