.DEFAULT_GOAL=all

.PHONY: all build up down
all: build up

setup:
	mkdir -p tmp/

build: setup
	docker compose build 

up:build
	docker compose up

down:
	docker compose down

clean:
	@echo "=> clean ... "
	rm -rf tmp/