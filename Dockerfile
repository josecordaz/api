FROM golang:1.13.2-alpine3.10

EXPOSE 8283

WORKDIR $GOPATH/src/josecordaz/api/

COPY main.go .

RUN go build -o /go/bin/api

ENTRYPOINT ["/go/bin/api"]