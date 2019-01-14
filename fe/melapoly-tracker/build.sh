#!/bin/bash

export APP_NAME=fe

# Build frontend.
npm run build

# Build docker image.
TAG=melapoly_tracker_$APP_NAME:latest
docker build -f Dockerfile -t $TAG .
