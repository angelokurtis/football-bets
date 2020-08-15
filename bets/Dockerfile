FROM golang:1.15.0-alpine AS build

ADD . $GOPATH/src/github.com/angelokurtis/football-bets/bets
WORKDIR $GOPATH/src/github.com/angelokurtis/football-bets/bets

RUN go get -d -v
RUN go build -o /go/bin/bets

FROM scratch

COPY --from=build /go/bin/bets /go/bin/bets
EXPOSE 8080

ENTRYPOINT ["/go/bin/bets"]

