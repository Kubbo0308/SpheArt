# 変数設定
FRONTEND_CONTAINER := frontend
BACKEND_CONTAINER := backend
MYSQL_CONTAINER := mysql
POSTGRESQL_CONTAINER := postgresql

# コマンド
.PHONY: build
build:
	docker compose build --no-cache

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: restart
restart:
	docker compose up --force-recreate --build --abort-on-container-exit

.PHONY: up-frontend
up-frontend:
	docker compose up -d ${FRONTEND_CONTAINER}

.PHONY: up-backend
up-backend:
	docker compose up -d ${BACKEND_CONTAINER}

generate-mock-data-for-test: up-backend
	.scripts/generate-mock-data-for-test.sh ${BACKEND_CONTAINER}
