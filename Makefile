.PHONY: build build-frontend build-backend start stop

build: build-frontend build-backend

build-frontend:
	docker-compose build web
	
build-backend:
	docker-compose build api

start: build
	docker-compose up -d --remove-orphans 

stop:
	docker-compose down --remove-orphans