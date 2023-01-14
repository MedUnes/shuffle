COMPOSE_COMMAND := $$(which docker-compose ||  echo "docker compose")
DOCKER_DIR=./.shuffle
COMPILE_DIR="${DOCKER_DIR}/data/build"
include ${DOCKER_DIR}/.env

DOCKER_BUILDKIT=1
COMPOSE_DOCKER_CLI_BUILD=1

build:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml build --no-cache
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml up -d
up:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml up -d
top:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml stop
login:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml exec go bash
down:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml down -v
status:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml ps
logs:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml logs -f
stop:
	${COMPOSE_COMMAND} --env-file=${DOCKER_DIR}/.env -f ${DOCKER_DIR}/docker-compose.yml stop
