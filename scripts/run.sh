#!/bin/bash

##################################################
#
# set GOPATH  GOROOT before running this script
#
##################################################




echo "...removing project..."
rm -rf goshell

echo "...show delete result..."
ls

echo "...cloning project..."
git clone https://github.com/leighyu1983/go-shell.git

echo "...mv go-shell goshell..."
mv go-shell goshell

echo "...show mv result..."
ls

echo "...move rpms to the same directory as executable binary..."
cp ${GOPATH}/scripts/rpms ${GOPATH}/src/applications/

echo "...change dir..."
cd ${GOPATH}/src/applications

echo "...set linux temp env..."
CGO_ENABLED=0 GOOS=linux

echo "...building golang..."
go build -a -installsuffix cgo  -tags netgo -o go-shell
