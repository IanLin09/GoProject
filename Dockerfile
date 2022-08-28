FROM golang

RUN     mkdir -p /app
WORKDIR /app

COPY    . .

RUN     export GO111MODULE=ON
RUN     go mod download
RUN     go build -o app

