#!/bin/bash

export OTEL_AGENT_VERSION=1.19.2

curl -Lo opentelemetry-javaagent-${OTEL_AGENT_VERSION}.jar  https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v${OTEL_AGENT_VERSION}/opentelemetry-javaagent.jar
