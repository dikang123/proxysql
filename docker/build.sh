#!/bin/bash

#set -e
set -x

readonly CNAME=proxysql_master
readonly IMGTAG=1.3

docker stop $CNAME
docker rm $CNAME

export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64

go build -o $CNAME  ../main.go

docker build -t $CNAME:$IMGTAG .

rm -f $CNAME

docker run --name $CNAME -p 3333:3333 -d $CNAME:$IMGTAG
