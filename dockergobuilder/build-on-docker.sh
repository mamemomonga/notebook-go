#!/bin/bash
set -eux

IMAGE_NAME=docker-go-build

make revision
docker build -t $IMAGE_NAME .
exec docker run --rm $IMAGE_NAME tar cC /app dist | tar xvp

