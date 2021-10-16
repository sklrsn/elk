.DEFAULT_GOAL=all

.PHONY: all build up down clean bootstrap logstash monitor elasticsearch kibana
all: build up

setup:
	mkdir -p tmp/

build: setup
	docker compose build 

rm: down clean

down:
	docker compose down

up: build
	docker compose up
	
clean:
	@echo "=> Removing ... Please wait"
	@rm -rf tmp/
	@rm -rf vendor/
	@docker system prune -f

bootstrap: setup
	@docker compose up bootstrap kafdrop

logstash: setup
	@docker compose build logstash
	@docker compose up logstash

elasticsearch:
	@docker compose up elasticsearch

kibana:
	@docker compose up kibana