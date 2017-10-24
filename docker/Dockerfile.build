FROM golang:1.9

WORKDIR /go/src

RUN \
	git clone https://github.com/taylor840326/proxysql-master.git && \
	cd proxysql-master && \
	go get -u github.com/go-sql-driver/mysql && \
	go get -u github.com/gin-gonic/gin && \
	go build main.go && \
	cp main /usr/local/bin/main

CMD  /usr/local/bin/main
