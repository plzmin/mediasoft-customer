build:
	docker compose up --build


migrate-up:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/customer?sslmode=disable up


migrate-down:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/customer?sslmode=disable down
