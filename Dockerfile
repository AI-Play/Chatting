FROM golang:1.17

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go mod download

WORKDIR /app/brandnew

EXPOSE 5000

RUN go build main.go

CMD ["/app/main.exe"]