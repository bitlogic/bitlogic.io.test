FROM golang:alpine
RUN apk --no-cache add git && \
    go get github.com/DATA-DOG/godog/cmd/godog
WORKDIR /go/src/github.com/bitlogic/bitlogic.io.test/
ENTRYPOINT ["/go/src/github.com/bitlogic/bitlogic.io.test/drone-entrypoint.sh"]
ADD . /go/src/github.com/bitlogic/bitlogic.io.test/
RUN go get -t ./...

