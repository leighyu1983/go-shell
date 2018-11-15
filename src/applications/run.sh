#!/bin/bash

echo "...removing project..."
rm -rf go-shell

echo "...cloning project..."
git clone https://github.com/leighyu1983/go-shell.git

echo "...set linux temp env..."
CGO_ENABLED=0 GOOS=linux

echo "...building golang..."
go build -a -installsuffix cgo  -tags netgo -o go-shell



