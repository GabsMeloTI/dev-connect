# migrate create -ext sql -dir db/migration -seq {{name}}
# include
include .env

run:
	go run main.go api

sqlc:
	sqlc generate

up:
	docker-compose up -d

down:
	docker-compose down