FROM golang:1.13.2-alpine3.10

EXPOSE 8283

WORKDIR $GOPATH/go/bin/

COPY api .

# RUN go build -o /go/bin/api

ENTRYPOINT ["/go/bin/api"]