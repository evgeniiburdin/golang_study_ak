d-comp-u:
	docker compose up --build
.PHONY: d-comp-u

d-comp-d:
	docker compose down --remove-orphans
.PHONY: d-comp-d

swag:
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

run: swag
	go mod tidy && go mod download && \
	go run ./cmd/app
.PHONY: run
