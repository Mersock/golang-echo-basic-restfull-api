FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go mod download

RUN go get -v github.com/cosmtrek/air

ENTRYPOINT [ "air" ,"-c",".air.toml"]
