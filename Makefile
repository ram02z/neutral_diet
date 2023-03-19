DB_USER ?= postgres
DB_PASS ?= postgres
DB_HOST ?= localhost
DB_NAME ?= postgres
DB_PORT ?= 5432

.PHONY: all
all:
	$(MAKE) gen-proto
	$(MAKE) gen-models

.PHONY: gen-proto
gen-proto:
	buf mod update internal/proto
	buf lint
	buf format -w
	buf generate

.PHONY: gen-models
gen-models:
	sqlc compile --experimental
	sqlc generate --experimental

.PHONY: migrate-up
migrate-up:
	docker run -v $(shell pwd)/db/migrations:/migrations \
	  --rm \
	  --network host \
	  migrate/migrate \
	  -path=/migrations/ \
	  -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)\?sslmode=disable \
	  up

.PHONY: migrate-down
migrate-down:
	docker run -v $(shell pwd)/db/migrations:/migrations \
	  --rm \
	  --network host \
	  migrate/migrate \
	  -path=/migrations/ \
	  -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)\?sslmode=disable \
	  down -all
