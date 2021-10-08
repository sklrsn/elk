.DEFAULT_GOAL=all

.PHONY: all build up down
all: build up

build:
	docker compose build 

up:build
	docker compose up

down:
	docker compose down