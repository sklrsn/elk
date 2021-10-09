FROM golang:latest

LABEL author="Kalai"

WORKDIR /go/src/github.com/sklrsn/elk

ADD . .

RUN go mod init github.com/sklrsn/elk

RUN go mod tidy

RUN go build -o /opt/elk/rabbitmq rabbitmq.go

RUN chmod +x /opt/elk/rabbitmq

ENTRYPOINT [ "/opt/elk/rabbitmq" ]