.DEFAULT_GOAL=all

.PHONY: all build up down clean bootstrap logstash monitor
all: build up

setup:
	mkdir -p tmp/

build: setup
	docker compose build 

rm: down clean

down:
	docker compose down

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
	@docker compose up elasticsearch kibana