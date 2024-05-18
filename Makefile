include .env

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

.PHONY: up-mysql
up-mysql:
	docker compose up -d ${MYSQL_CONTAINER}

.PHONY: up-postgresql
up-postgresql:
	docker compose up -d ${POSTGRESQL_CONTAINER}

backend-container: up-backend
	docker compose exec ${BACKEND_CONTAINER} bash

mysql-container: up-mysql
	docker compose exec ${MYSQL_CONTAINER} mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}

postgresql-container: up-postgresql
	docker compose exec ${POSTGRESQL_CONTAINER} sh -c "PGPASSWORD=${POSTGRES_PASSWORD} psql -U ${POSTGRES_USER} -d ${POSTGRES_DATABASE}"

generate-mock-data-for-test: up-backend
	./scripts/generate-mock-data-for-test.sh ${BACKEND_CONTAINER}
