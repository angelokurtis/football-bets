FROM openjdk:14-alpine
COPY build/libs/championships-*-all.jar championships.jar
EXPOSE 8080
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx128m", "-jar", "championships.jar"]