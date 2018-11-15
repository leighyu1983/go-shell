#!/bin/bash

echo "...set linux temp env..."
CGO_ENABLED=0 GOOS=linux

echo "...building golang..."
go build -a -installsuffix cgo  -tags netgo -o go-shell



