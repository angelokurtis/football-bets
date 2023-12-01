FROM golang:1.21.4 as builder

WORKDIR /go/src/app
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go
COPY internal/ internal/
COPY docs/ docs/

RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/app main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /go/bin/app .

USER 65532:65532
EXPOSE 8081/tcp

CMD ["/app"]
