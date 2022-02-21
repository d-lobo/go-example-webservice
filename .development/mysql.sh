#!/bin/sh

cd "$(dirname "$0")"

docker run --name test-mysql --rm \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=root \
	-e DB_USER=user \
	-e DB_USER_PASSWORD=secret \
	-e DB_NAME=test \
	mysql &

terminate() {
	docker stop test-mysqls
}

trap terminate 2 15
wait