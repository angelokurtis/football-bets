FROM golang:1.15-buster AS builder

WORKDIR $GOPATH/src/bets
COPY . .

RUN go get -d -v
RUN go build -o /bin/bets

FROM gcr.io/distroless/base

COPY --from=builder /bin/bets /bin/bets
EXPOSE 9090

CMD ["/bin/bets"]

