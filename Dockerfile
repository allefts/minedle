FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get -t .
RUN go build -o main .

CMD ["/app/main"]