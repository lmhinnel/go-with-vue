FROM golang:latest

RUN go get -u github.com/jinzhu/gorm

COPY . /go/src/github.com/kansuke231/go-with-vue/api
WORKDIR /go/src/github.com/kansuke231/go-with-vue/api

RUN go get ./
RUN go build