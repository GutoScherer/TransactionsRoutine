#!make

configure:
	cp .env.example .env

start:
	docker-compose up

stop:
	docker-compose stop

restart:
	docker-compose restart

reset-database:
	docker-compose stop db && docker-compose up --renew-anon-volumes

test:
	docker exec TransactionsRoutineApp bash -c 'go test -race -cover ./...'