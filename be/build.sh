#!/bin/bash

export APP_NAME=be
export APP_REPO=github.com/alex-ant/melapoly-tracker/$APP_NAME
TAG=melapoly_tracker_be:latest

# build statically-linked binary
docker run --name "$APP_NAME"_build -v "$PWD":/go/src/$APP_REPO -w /go/src/$APP_REPO -e CGO_ENABLED=0 -d golang:1.10 sleep infinity
docker exec "$APP_NAME"_build go build --ldflags '-extldflags "-static"'

docker stop "$APP_NAME"_build && docker rm "$APP_NAME"_build

if [ $? != 0 ]; then
  echo "failed to build the project"
  exit 1
fi

echo "build successful"

# build Docker image
docker build -f Dockerfile -t $TAG .

# remove the binary
rm be
