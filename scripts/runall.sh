#!/bin/bash

set -e

export SCRIPT_PATH=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
export OPENTELEMETRY_JAVAAGENT_PATH="${SCRIPT_PATH}/../opentelemetry-javaagent-1.19.2.jar"
export OTEL_TRACES_EXPORTER=otlp
export OTEL_METRICS_EXPORTER=none
export OTEL_LOGS_EXPORTER=none
export JAVA_TOOL_OPTIONS="-javaagent:${OPENTELEMETRY_JAVAAGENT_PATH}"
export MATCHES_URL=http://localhost:8083
export TEAMS_URL=http://localhost:8084
export CHAMPIONSHIPS_URL=http://localhost:8082
export MATCHES_JSON_PATH="${SCRIPT_PATH}/../matches/matches.json"
export RELATIONSHIPS_JSON_PATH="${SCRIPT_PATH}/../matches/relationships.json"
export TEAMS_JSON_PATH="${SCRIPT_PATH}/../teams/teams.json"
export CHAMPIONSHIPS_JSON_PATH="${SCRIPT_PATH}/..//championships/championships.json"

function runasync() {
  echo "running $1"
 ( cd $(realpath $1) ; OTEL_SERVICE_NAME=$1 nohup ./mvnw spring-boot:run >/dev/null 2>&1 & )
}

function runsync() {
  echo "running $1"
 ( cd $(realpath $1) ; OTEL_SERVICE_NAME=$1 ./mvnw spring-boot:run)
}

pkill -f maven-wrapper.jar || true
runasync championships
runasync matches
runasync teams
runsync bets
