up:
	docker-compose up

swag:
	docker-compose up swagger-ui

mysql:
	docker-compose up mysql

down:
	docker-compose down

stop:
	docker-compose stop

del:
	docker system prune -a
