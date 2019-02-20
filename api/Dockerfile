FROM golang:latest

RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/patrickmn/go-cache

COPY . /go/src/github.com/kansuke231/go-with-vue/api
WORKDIR /go/src/github.com/kansuke231/go-with-vue/api

RUN go get ./
RUN go build