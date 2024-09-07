.PHONY: build build-frontend build-backend start

build: build-frontend build-backend

build-frontend:
	docker-compose build web
	
build-backend:
	docker-compose build api

start: build
	docker-compose up --remove-orphans