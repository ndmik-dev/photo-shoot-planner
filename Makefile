run:
	go run ./cmd/api

db-up:
	docker compose up -d

db-down:
	docker compose down

migrate-up:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5433/photo_planner?sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5433/photo_planner?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test ./...