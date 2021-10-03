create:
	docker-compose -f docker-compose-dev.yaml up --build

stop:
	docker-compose stop

start:
	docker-compose start
	docker-compose logs -f --since 0m

.PHONY:create stop start