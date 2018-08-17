#!/bin/bash

docker run --rm -it -v ${PWD}:/go golang:stretch env GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -ldflags -s -a -installsuffix cgo main.go

CGO_ENABLE=0 go build -ldflags -s -a -installsuffix cgo main.go