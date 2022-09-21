FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

ENV MATCHES_JSON_PATH=matches.json
ENV RELATIONSHIPS_JSON_PATH=relationships.json

ADD matches.json .
ADD relationships.json .
COPY --from=build target/matches-1.0.0.jar .
EXPOSE 8083 9093

CMD ["java", "-jar", "matches-1.0.0.jar"]
