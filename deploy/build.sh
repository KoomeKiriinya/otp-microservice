#!/bin/bash

docker build --build-arg REDIS_DSN=${REDIS_DSN} \
 --build-arg REDIS_PASSWORD=${REDIS_PASSWORD} \
 -t ${IMAGE_NAME}:${TAG} .