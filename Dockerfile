FROM golang:latest

LABEL maintainer="Dr.Vishwanath Rao"


WORKDIR /app


COPY go.mod ./


RUN go mod download

COPY . .

RUN go build -o main ./go-docker

EXPOSE 4040


CMD [ "./main" ]

