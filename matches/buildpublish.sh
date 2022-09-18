#!/bin/bash

export DOCKER_BUILDKIT=1

docker build -t kurtis/matches:1.0.0-java-spring .
docker push kurtis/matches:1.0.0-java-spring
