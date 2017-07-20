#!/bin/bash

export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64


go build -o proxysql-master ../main.go

docker build -t pm .
