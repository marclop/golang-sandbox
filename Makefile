COMPONENT ?= golang
VERSION ?= None
ENV ?= staging
SERVICE ?= api
CODE_CONTAINER := ${golang}

all: dev
ps: status
restart: nodev dev

dev:
	@docker-compose -p ${COMPONENT} -f ops/docker/docker-compose.yml up -d
	##@./ops/scripts/hooks.sh

enter:
	@./ops/scripts/enter.sh ${COMPONENT}

kill:
	@docker-compose -p ${COMPONENT} -f ops/docker/docker-compose.yml kill

nodev:
	@docker-compose -p ${COMPONENT} -f ops/docker/docker-compose.yml kill
	@docker-compose -p ${COMPONENT} -f ops/docker/docker-compose.yml rm -f

status:
	@docker-compose -p ${COMPONENT} -f ops/docker/docker-compose.yml ps
