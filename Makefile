.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t hmochizuki/gotodo:${DOCKER_TAG} \
			    --target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Run docker container UP with hot reload
	docker compose up -d

down: ## Stop docker container
	docker compose down

logs: ## Run docker container with logs
	docker compose logs -f

ps: ## Show docker container status
	docker compose ps

test: ## Run tests
	go test -v ./...

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'