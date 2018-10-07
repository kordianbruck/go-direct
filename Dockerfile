FROM golang:alpine as builder
COPY ./src $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/
RUN go get -d -v
RUN go build -o /go/bin/hello

FROM scratch
COPY --from=builder /go/bin/hello /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]