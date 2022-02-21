## Intent

This is an example implementation of a tiny microservice written in Go,
using `mongo` as storage and `gin` as web framework.
This app is not meant to be deployed anywhere,
thus ignoring typical patterns like reading secrets from env-properties or external vault tools

## How to build

* `go mod tidy && go build`

## How to run

requires a local mongo db instance (you may use provided script `.development/mongo.sh`)

* `go install` & execute binary or `go run .`

## Example requests

requires `curl` & `jq`

* getting all albums
  * `curl localhost:8080/albums | jq .`
* inserting an album
  * `curl -H 'Content-Type: application/json' localhost:8080 -d '{"title":"exampleTitle", "artist":"exampleArtist", "price":9.99}'`
* getting a specific album
  * `curl localhost:8080/albums/<id> | jq .`