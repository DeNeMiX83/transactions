include deploy/.env

compose-build:
	docker compose -f ./deploy/docker-compose.yml --env-file deploy/.env build

compose-up:
	docker compose -f ./deploy/docker-compose.yml --env-file deploy/.env up

docker-rm-volume:
	docker volume rm -f ${PROJECT_NAME}_database_data

# docker-rabbitmq-run:
# 	docker run -d \
# 	--hostname rabbitmq \
# 	--name rabbitmq \
# 	-p 5672:5672 -p 15672:15672 \
# 	-e RABBITMQ_DEFAULT_USER=${RABBITMQ_USER} \
# 	-e RABBITMQ_DEFAULT_PASS=${RABBITMQ_PASSWORD} \
# 	rabbitmq:3-management

# docker-rabbitmq-stop:
# 	docker stop rabbitmq

# docker-rabbitmq-start:
# 	docker start rabbitmq