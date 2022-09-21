FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

ENV CHAMPIONSHIPS_JSON_PATH=championships.json

ADD championships.json .
COPY --from=build target/championships-1.0.0.jar .
EXPOSE 8082 9092

CMD ["java", "-jar", "championships-1.0.0.jar"]
