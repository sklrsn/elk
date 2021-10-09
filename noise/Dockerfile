FROM golang:latest

WORKDIR /go/src/github.com/sklrsn/elk/

ADD . /go/src/github.com/sklrsn/elk/

RUN go build -o bin/noise noise.go

RUN chmod +x bin/noise

ENTRYPOINT [ "bin/noise" ] 