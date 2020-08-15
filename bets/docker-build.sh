#!/bin/sh

set -e

docker build . -t kurtis/bets:v1
echo
echo
echo "To run the docker container execute:"
echo "    $ docker run -p 9090:9090 kurtis/bets:v1"
