REPO := $(shell pwd)

redis-run:
	docker run -d --name localNode1 -p 6389:6379  -d redis:latest redis-server --appendonly yes --requirepass "123456"

rabbitmq:
    docker run -d --hostname rabbit-node1 --name rabbit-node1 -p 5672:5672 -p 15672:15672 rabbitmq:management


mysql:
    docker run -d --name mysql1 -e MYSQL_ROOT_PASSWORD=zxcasdqwe!! -p 3306:3306 mysql