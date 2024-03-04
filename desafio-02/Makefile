STACK_NAME=lab-desafio-02
CONTAINER_SERVICE_A=service-a

start:
	docker-compose -p ${STACK_NAME} up -d

build:
	docker-compose -p ${STACK_NAME} up --build -d

stop:
	docker-compose -p ${STACK_NAME} stop

restart: stop start

clean:
	docker-compose -p ${STACK_NAME} down

logs:
	docker-compose -p ${STACK_NAME} logs -f


log-service-a:
	docker logs -f ${CONTAINER_SERVICE_A}

ps:
	docker-compose -p ${STACK_NAME} ps
