FROM golang:1.17

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go mod download

WORKDIR /app/brandnew

RUN go build main.go

EXPOSE 5000

CMD ["nohup", "./main", "&"]