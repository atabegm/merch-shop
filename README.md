# Merch Shop

Backend-сервис магазина мерча на Go.

Пользователь получает 1000 монет после регистрации, может покупать мерч, переводить монеты другим пользователям и смотреть баланс, инвентарь и историю переводов.

После покупки сервис отправляет событие в Kafka. Отдельный consumer получает событие и отправляет email через SMTP. Для локальной проверки почты используется Mailpit.

## Стек

- Go
- PostgreSQL
- pgx
- JWT
- Apache Kafka
- Mailpit
- Docker Compose
- gomock
- testify

## API

### Авторизация

`POST /api/auth`

При первой авторизации пользователь создаётся автоматически.  
Если пользователь уже существует и передал новый email, email обновляется.

### Информация о пользователе

`GET /api/info`

Возвращает баланс, инвентарь и историю переводов.

### Перевод монет

`POST /api/send`

Пример:

```json
{
  "toUser": "ivan",
  "amount": 100
}
```

### Покупка мерча

`GET /api/buy/{item}`

После успешной покупки монеты списываются, покупка сохраняется в БД и событие отправляется в Kafka.

Consumer получает событие и отправляет письмо через SMTP.

## Запуск

Создать `.env` на основе `.env_example`.

Поднять PostgreSQL:

```bash
docker compose up -d db
```

Применить миграции:

```bash
make migrate-up
```

Запустить проект:

```bash
docker compose up -d --build
```

API: `http://localhost:8080`

Mailpit: `http://localhost:8025`

## Тесты

В проекте есть unit-тесты сервисного и API-слоя, а также интеграционные тесты с PostgreSQL и Mailpit.

```bash
go test -v ./...
```

## Структура

```text
cmd/app              - API
cmd/consumer         - Kafka consumer
internal/api         - HTTP handlers
internal/service     - бизнес-логика
internal/repository  - работа с PostgreSQL
internal/kafka       - producer и consumer
internal/model       - модели
migrations           - миграции
```
