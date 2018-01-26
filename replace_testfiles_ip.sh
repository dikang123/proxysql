#!/bin/bash

set -xeuo pipefail

readonly OLDIP=$1
readonly NEWIP=$2
readonly DIR=`pwd`

for testfile in `ls $DIR/*_test.go`;
do
	sed -i "" "s/$OLDIP/$NEWIP/g" $testfile
done
