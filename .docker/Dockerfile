FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/go-dep-test

WORKDIR /go/src/github.com/Sharykhin/go-dep-test

RUN go get . && \
    go get github.com/canthefason/go-watcher && \
    go install github.com/canthefason/go-watcher/cmd/watcher && \

EXPOSE 5000

ENTRYPOINT watcher
