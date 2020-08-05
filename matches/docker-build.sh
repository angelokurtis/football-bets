#!/bin/sh

set -e

./mvnw clean package -Dquarkus.container-image.build=true -DskipTests
echo
echo
echo "To run the docker container execute:"
echo "    $ docker run --rm -v \$(pwd):/usr/src/data -p 8080:8080 kurtis/matches:v1"
