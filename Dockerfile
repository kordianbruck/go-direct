FROM golang:alpine as builder

COPY ./src $GOPATH/src/
WORKDIR $GOPATH/src/

RUN go build -o /go-direct
RUN chmod +x /go-direct

FROM alpine:latest
COPY --from=builder /go-direct /go-direct

CMD ["/go-direct"]
