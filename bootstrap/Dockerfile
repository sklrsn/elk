FROM golang:latest

LABEL author="Kalai"

WORKDIR /go/src/github.com/sklrsn/elk

ADD . .

RUN go mod init github.com/sklrsn/elk

RUN go mod tidy

RUN go build -o /opt/elk/bootstrap bootstrap.go

RUN chmod +x /opt/elk/bootstrap

ADD entrypoint.sh /opt/elk/

RUN chmod +x /opt/elk/entrypoint.sh

CMD [ "/opt/elk/entrypoint.sh" ]