FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

ENV OTEL_AGENT_VERSION=1.27.0
ENV JAVA_TOOL_OPTIONS="-javaagent:opentelemetry-javaagent.jar"

ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v${OTEL_AGENT_VERSION}/opentelemetry-javaagent.jar .
COPY --from=build target/bets-1.0.0.jar .
EXPOSE 8081 9091

CMD ["java", "-jar", "bets-1.0.0.jar"]
