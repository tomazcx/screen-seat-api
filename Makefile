help: ## Prints available commands
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[.a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test: ## Run all application's tests
	go test ./...

MIGRATION_NAME =
new-migration: ## Create new migration
	migrate create -ext=sql -dir=./internal/infra/database/migrations $(MIGRATION_NAME)

migrate-up:  ## Run migrations up
	migrate -source file:./internal/infra/database/migrations/ -database "postgres://root:root@db:5432/screen-seat-api" up

migrate-down: ## Run migrations down 1
	migrate -source file:./internal/infra/database/migrations/ -database "postgres://root:root@db:5432/screen-seat-api" down 1

up: ## Run the application containers and its dependencies
	docker-compose down --remove-orphans
	docker-compose up

build: ## Run and build the containers
	docker-compose down --remove-orphans --rmi all
	docker-compose up --build

down: ## Stop and destroy the containers
	docker-compose down --remove-orphans

down-rmi: ## Stop and destroy the containers and images
	docker-compose down --remove-orphans --rmi all
