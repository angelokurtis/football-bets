FROM golang:1.15.0-alpine AS builder

WORKDIR $GOPATH/src/bets
COPY . .

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/bets

FROM alpine

COPY --from=builder /bin/bets /bin/bets
EXPOSE 9090

CMD ["/bin/bets"]

