# Photo Shoot Planner

Small Go REST API for planning photo shoots.  
The project uses:

- Go
- PostgreSQL
- `sqlc` for query generation
- `goose` for migrations

## Features

- health check endpoint
- create shoot endpoint
- list shoots endpoint

## API

- `GET /health`
- `POST /api/v1/shoots/`
- `GET /api/v1/shoots/`

Default app address: `http://localhost:8080`

## Requirements

- Go installed
- Docker and Docker Compose
- `goose` installed
- `sqlc` installed

## Environment

Copy `.env.example` to `.env` if you want local overrides.

Default values:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5433
DB_NAME=photo_planner
DB_USER=postgres
DB_PASSWORD=postgres
DB_SSLMODE=disable
```

## Start The App

1. Start PostgreSQL:

```bash
make db-up
```

2. Run database migrations:

```bash
make migrate-up
```

3. Start the API:

```bash
make run
```

The app will connect to PostgreSQL on `localhost:5433` and start on `localhost:8080`.

## Useful Commands

```bash
make db-down
make migrate-down
make sqlc
make test
```
