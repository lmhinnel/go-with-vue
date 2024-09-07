FROM golang:latest

COPY . /go/src/github.com/kansuke231/go-with-vue/api
WORKDIR /go/src/github.com/kansuke231/go-with-vue/api

RUN go mod tidy
RUN go test
RUN go install