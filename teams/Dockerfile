FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

ENV DATA_DIR=teams.json

ADD teams.json .
COPY --from=build target/teams-1.0.0.jar .
EXPOSE 8084 9094

CMD ["teams-1.0.0.jar"]
