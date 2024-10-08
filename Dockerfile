FROM golang:latest as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /main

EXPOSE 5000

CMD ["/main"]