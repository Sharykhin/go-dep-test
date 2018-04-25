FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/go-dep-test

WORKDIR /go/src/github.com/Sharykhin/go-dep-test

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure && \
    go build -o /go/bin/go-dep-test .

EXPOSE 5000

ENTRYPOINT /go/bin/go-dep-test