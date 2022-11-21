FROM golang:1.18.1-bullseye

RUN apt-get update && apt install -y netcat

WORKDIR /golang-backend

COPY go*.* ./

RUN go get ./...

COPY . .

RUN chmod +x wait-for

CMD go run main.go