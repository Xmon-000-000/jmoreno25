FROM golang:1.26.2-alpine3.23

WORKDIR /app

COPY main.go .

RUN go build -o server main.go

EXPOSE 9090

CMD ["/app/server"]
