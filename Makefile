up:
	docker-compose up

swag:
	docker-compose up swagger-ui

mysql:
	docker-compose exec mysql mysql -u root -p ca-tech-dojo

down:
	docker-compose down

stop:
	docker-compose stop

del:
	docker system prune -a
