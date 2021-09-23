FROM golang:1.17

WORKDIR /go/src/github.com/winadiw/go-marvel-api

COPY . .

RUN go mod tidy

RUN go build -o go-marvel-api

EXPOSE 8080
CMD ./go-marvel-api