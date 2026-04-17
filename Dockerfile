FROM golang:1.26-alpine3.22 as builder

COPY ./src $GOPATH/src/
COPY go.mod $GOPATH/src/go.mod
WORKDIR $GOPATH/src/

RUN go build -o /go-direct
RUN chmod +x /go-direct

FROM alpine:latest
COPY --from=builder /go-direct /go-direct

CMD ["/go-direct"]
