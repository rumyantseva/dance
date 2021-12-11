all: fmt dance

help:                                  ## Display this help message
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

env-up: env-up-detach                  ## Start development environment
	docker-compose logs --follow

env-up-detach:
	docker-compose up --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes --detach

env-pull:
	docker-compose pull --include-deps --quiet

env-down:                              ## Stop development environment
	docker-compose down --remove-orphans

init:                                  ## Install development tools
	go mod tidy
	cd tools && go mod tidy
	go mod verify
	cd tools && go generate -tags=tools -x

fmt: bin/gofumpt                       ## Format code
    # skip submodules
	bin/gofumpt -w ./cmd/ ./internal/

dance:                                 ## Run all integration tests
	cd tests && go run ../cmd/dance

lint: bin/golangci-lint                ## Run linters
	bin/golangci-lint run --config=.golangci-required.yml
	bin/golangci-lint run --config=.golangci.yml
	bin/go-consistent -v -pedantic ./cmd/... ./internal/...

psql:                                  ## Run psql
	docker-compose exec postgres psql -U postgres -d ferretdb

mongosh:                               ## Run mongosh
	docker-compose exec mongodb mongosh mongodb://ferretdb:27017/ \
		--verbose --eval 'disableTelemetry()' --shell

mongo:                                 ## Run (legacy) mongo shell
	docker-compose exec mongodb mongo mongodb://ferretdb:27017/ \
		--verbose

bin/golangci-lint:
	$(MAKE) init