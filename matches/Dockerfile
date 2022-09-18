FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

ENV DATA_DIR=matches.json

ADD matches.json .
COPY --from=build target/matches-1.0.0.jar .
EXPOSE 8083 9093

CMD ["matches-1.0.0.jar"]
