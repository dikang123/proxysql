#!/bin/bash

set -e
set -x

docker stop pm
docker rm pm

export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64

go build -o main  ../main.go

docker build -t pm .

rm -f main

docker run --name pm -p 3333:3333 -d pm 
