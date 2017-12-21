FROM golang:1.9
MAINTAINER Lei Tian

WORKDIR /go/src

RUN \
    mkdir -p github.com/imSQL && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/gin-gonic/gin && \
    git clone https://github.com/imSQL/proxysql-master.git github.com/imSQL/proxysql-master && \
    go install proxysql-master


ENTRYPOINT ["/go/bin/proxysql-master"]
