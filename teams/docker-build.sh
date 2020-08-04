#!/bin/sh

set -e

docker build . -t kurtis/teams:v1
echo
echo
echo "To run the docker container execute:"
echo "    $ docker run -p 3000:3000 kurtis/teams:v1"
