FROM golang:latest

COPY . /go/src/github.com/kansuke231/go-with-vue/newsfeeder
WORKDIR /go/src/github.com/kansuke231/go-with-vue/newsfeeder

RUN go mod tidy
RUN go install