FROM golang:latest

RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/mmcdole/gofeed

COPY . /go/src/github.com/kansuke231/go-with-vue/newsfeeder
WORKDIR /go/src/github.com/kansuke231/go-with-vue/newsfeeder

RUN go get ./
RUN go install