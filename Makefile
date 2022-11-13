.PHONY: all
all:
	$(MAKE) gen-proto
	$(MAKE) gen-models

.PHONY: gen-proto
gen-proto:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf mod update internal/proto
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf format -w
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf generate

.PHONY: gen-models
gen-models:
	docker run --rm --volume "$(shell pwd):/src" --workdir /src kjconroy/sqlc compile
	docker run --rm --volume "$(shell pwd):/src" --workdir /src kjconroy/sqlc generate

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
