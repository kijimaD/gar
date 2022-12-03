.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## デプロイ用のイメージをビルドする
	docker build -t kijimad/gotodo:${DOCKER_TAG} \
	--target deploy ./

build-local: ## ローカル開発用のイメージをビルドする
	docker-compose build --no-cache

up: ## Do docker compose up
	docker-compose up -d

down: ## Do docker compose down
	docker-compose down

logs: ## Tail docker compose logs
	docker-compose logs -f

ps: ## Check container status
	docker-compose ps

gen: ## go generateを実行する
	go generate ./...

test: ## テストを実行する
	go test ./... -v -race -shuffle=on

testGH: ## GitHubへのアクセスが発生するテストを実行する
	go test ./... -v -tags gh

help: ## ヘルプを表示する
	@echo -e "──────● █▀▀ ▄▀█ █▀█ ●──────\n        █▄█ █▀█ █▀▄        "
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
