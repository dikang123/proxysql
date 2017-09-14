#!/bin/bash

set -e
set -x

export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64

GW=`docker network inspect b9d9659ce86c |grep Gateway |awk -F':' '{print $2}'|awk -F'"' '{print $2}'`

go build -o main  ../main.go

docker build -t pm .

rm -f main

docker run -p 3333:3333 -d pm 
