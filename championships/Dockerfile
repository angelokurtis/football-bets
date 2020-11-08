FROM oracle/graalvm-ce:20.1.0-java8 as graalvm
RUN gu install native-image

COPY . /home/app/championships
WORKDIR /home/app/championships

RUN native-image --no-server -cp build/libs/championships-*-all.jar

FROM frolvlad/alpine-glibc
RUN apk update && apk add libstdc++
ENV DATA_DIR=/usr/src/data
ADD ./data.tar.xz $DATA_DIR/
EXPOSE 8080
COPY --from=graalvm /home/app/championships/championships /app/championships
ENTRYPOINT ["/app/championships"]
