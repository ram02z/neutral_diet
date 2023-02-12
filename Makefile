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
	  -database postgres://postgres:postgres@localhost:5432/postgres\?sslmode=disable \
	  up

.PHONY: migrate-down
migrate-down:
	docker run -v $(shell pwd)/db/migrations:/migrations \
	  --rm \
	  --network host \
	  migrate/migrate \
	  -path=/migrations/ \
	  -database postgres://postgres:postgres@localhost:5432/postgres\?sslmode=disable \
	  down -all
