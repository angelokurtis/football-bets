FROM amazoncorretto:18.0.2 as build

COPY .mvn .mvn
COPY mvnw .
COPY pom.xml .
COPY src src
RUN --mount=type=cache,target=/root/.m2 ./mvnw -B package

FROM amazoncorretto:18.0.2

COPY --from=build target/bets-1.0.0.jar .
EXPOSE 8081 9091

CMD ["bets-1.0.0.jar"]
