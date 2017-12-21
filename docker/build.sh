#!/bin/bash

set -xeuo pipefail

readonly HTTP_PROXY=http://172.18.0.2:7777/pac
readonly CNAME=proxysql_master
readonly IMGTAG=1.3

#docker stop $CNAME
#docker rm -f $CNAME

#docker build with proxy
#docker build -t $CNAME --build-arg HTTP_PROXY=$HTTP_PROXY .

#docker build without proxy
docker build -t $CNAME .
