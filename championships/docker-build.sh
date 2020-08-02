#!/bin/sh

set -e

./gradlew assemble
docker build . -t kurtis/championships:v1
echo
echo
echo "To run the docker container execute:"
echo "    $ docker run -p 8080:8080 kurtis/championships:v1"
