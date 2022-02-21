#!/bin/sh

MONGO_VERSION=${MONGO_VERSION:-5.0.6}
CONTAINER_NAME=mongo-test

[ $(id -u) -eq 0 ] || echo 'You may need to run this script as root or add your user to the docker group!' >&2
docker --help >/dev/null 2>/dev/null || (echo 'Please install docker ' >&2; false) || exit 1

terminatePid() {
	kill -15 "$1" 2>/dev/null
	for SEC in 1 1 3 5; do
		if ! ps $CONTAINER_PID >/dev/null 2>/dev/null; then
			return 0;
		fi
	done
	echo 'Killing process since it did not respond' >&2
	kill -9 "$1"
	return 1
}

terminateGracefully() {
	trap : 2 3 9 15
	echo 'Terminating gracefully ...'
	docker stop "$CONTAINER_NAME" >/dev/null
	terminatePid "$CONTAINER_PID"
	docker rm -f "$CONTAINER_NAME" >/dev/null
}

(docker rm -f "$CONTAINER_NAME" 2>/dev/null >/dev/null || true) &&
docker run --name "$CONTAINER_NAME" --rm \
	-p 27017:27017 \
	"mongo:$MONGODB_VERSION" &
CONTAINER_PID=$!
trap terminateGracefully 2 3 9 15
wait

