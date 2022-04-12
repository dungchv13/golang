FROM golang:1.18.0-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN mkdir -p /home/app

COPY . .
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]

