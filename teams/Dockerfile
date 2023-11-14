FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

ENV TEAMS_JSON_PATH=teams.json
ENV OTEL_AGENT_VERSION=1.31.0
ENV JAVA_TOOL_OPTIONS="-javaagent:opentelemetry-javaagent.jar"

ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v${OTEL_AGENT_VERSION}/opentelemetry-javaagent.jar .
ADD teams.json .
COPY --from=build target/teams-1.0.0.jar .
EXPOSE 8084 9094

CMD ["java", "-jar", "teams-1.0.0.jar"]
