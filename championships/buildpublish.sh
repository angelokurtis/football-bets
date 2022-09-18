#!/bin/bash

export DOCKER_BUILDKIT=1

docker build -t kurtis/championships:1.0.0-java-spring .
docker push kurtis/championships:1.0.0-java-spring
