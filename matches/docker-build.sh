#!/bin/sh

set -e

./mvnw package -Pnative -Dquarkus.native.container-build=true -DskipTests
docker build -f src/main/docker/Dockerfile.native -t kurtis/matches:v1 .

echo
echo
echo "To run the docker container execute:"
echo "    $ docker run --rm -p 8080:8080 kurtis/matches:v1"
