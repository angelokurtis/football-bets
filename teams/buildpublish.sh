#!/bin/bash

export DOCKER_BUILDKIT=1

docker build -t kurtis/teams:1.0.0-java-spring .
docker push kurtis/teams:1.0.0-java-spring
