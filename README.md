# Merch Shop

Backend-сервис магазина мерча на Go.

При первой авторизации пользователь получает 1000 монет. Монеты можно тратить на мерч или переводить другим пользователям. Также доступны просмотр баланса, инвентаря и истории переводов.

После покупки сервис публикует событие в Kafka. Отдельный consumer получает событие и отправляет email через SMTP. Для локальной проверки писем используется Mailpit.

## Стек

- Go
- PostgreSQL, pgx
- JWT
- Apache Kafka
- `avito-tech/go-transaction-manager`
- Docker, Docker Compose
- Mailpit (SMTP)
- gomock, testify

## API

### Авторизация

`POST /api/auth`

При первой авторизации пользователь создаётся автоматически и получает 1000 монет. Если пользователь уже существует и передал новый email, адрес обновляется.

### Информация о пользователе

`GET /api/info`

Возвращает баланс, инвентарь и историю переводов.

### Перевод монет

`POST /api/send`

Пример запроса:

```json
{
  "toUser": "ivan",
  "amount": 100
}
```

### Покупка мерча

`GET /api/buy/{item}`

После успешной покупки монеты списываются, покупка сохраняется в PostgreSQL, а событие публикуется в Kafka. Consumer обрабатывает его и отправляет письмо через SMTP.

Дополнительное описание API находится в файлах `schema.yaml` и `schema.json`.

## Транзакции

Для операций покупки мерча и перевода монет используется библиотека `avito-tech/go-transaction-manager` с драйвером `pgx/v5`. Она управляет транзакциями PostgreSQL, позволяя выполнить связанные изменения базы данных в рамках одной транзакции.

## Инструкция по запуску

Для запуска необходимы Docker с Docker Compose, Go, `make` и установленная утилита `golang-migrate` (команда `migrate`).

1. Склонировать репозиторий и перейти в директорию проекта:

   ```bash
   git clone https://github.com/atabegm/merch-shop.git
   cd merch-shop
   ```

2. Создать файл `.env` на основе `.env_example`:

   ```bash
   cp .env_example .env
   ```

3. Запустить PostgreSQL:

   ```bash
   docker compose up -d db
   ```

4. Применить миграции базы данных:

   ```bash
   make migrate-up
   ```

5. Собрать и запустить приложение и остальные сервисы:

   ```bash
   make docker-build
   ```

После запуска доступны:

- API: http://localhost:8080
- Mailpit: http://localhost:8025

При запуске миграций с компьютера PostgreSQL доступен по адресу `localhost:5433`. Контейнеры приложения используют внутренний адрес `db:5432`.

## Запуск тестов

В проекте есть unit-тесты API и сервисного слоя, а также интеграционные тесты с PostgreSQL и Mailpit.

Для запуска всех тестов:

1. Запустить PostgreSQL и Mailpit:

   ```bash
   docker compose up -d db mailpit
   ```

2. Применить миграции (если ещё не применены):

   ```bash
   make migrate-up
   ```

3. Запустить тесты:

   ```bash
   make test
   ```

Команда `make test` запускает `go test` с флагом `-race` для проверки гонок данных.

## Остановка

Для остановки и удаления контейнеров:

```bash
make docker-down
```

Остальные команды доступны в `Makefile`.

## Структура проекта

```text
cmd/app              — HTTP API
cmd/consumer         — Kafka consumer
internal/api         — HTTP-обработчики
internal/service     — бизнес-логика
internal/repository  — работа с PostgreSQL
internal/kafka       — Kafka producer и consumer
internal/model       — модели
migrations           — SQL-миграции
```