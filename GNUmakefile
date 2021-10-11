.DEFAULT_GOAL=all

.PHONY: all build up down clean
all: build up

setup:
	mkdir -p tmp/

build: setup
	docker compose build 

up:build
	docker compose up --force-recreate

rm: down clean

down:
	docker compose down

clean:
	@echo "=> Removing ... Please wait"
	@rm -rf tmp/
	@rm -rf vendor/
	@docker system prune -f