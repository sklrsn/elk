FROM golang:latest

WORKDIR /go/src/github.com/sklrsn/noise/

ADD . /go/src/github.com/sklrsn/noise/

RUN go build -o bin/noise noise.go

RUN chmod +x bin/noise

ENTRYPOINT [ "bin/noise" ] 