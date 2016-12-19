FROM golang:alpine
RUN apk --no-cache add git && \
    go get github.com/DATA-DOG/godog/cmd/godog
WORKDIR /go/src/github.com/bitlogic/bitlogic.io.test/
ADD . /go/src/github.com/bitlogic/bitlogic.io.test/
RUN go get -t ./...
ENTRYPOINT ["godog"]
