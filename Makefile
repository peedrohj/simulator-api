export UID:=$(shell id -u)

# Run dev containers
run-dev:
	docker-compose up $(args) app

# Build dev containers
build-dev:
	docker-compose build app