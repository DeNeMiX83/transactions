include deploy/.env

compose-build:
	docker compose -f ./deploy/docker-compose.yml --env-file deploy/.env build

compose-up:
	docker compose -f ./deploy/docker-compose.yml --env-file deploy/.env up

docker-rm-volume:
	docker volume rm -f ${PROJECT_NAME}_database_data