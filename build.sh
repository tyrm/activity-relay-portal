#!/bin/bash

go get github.com/gorilla/mux
go get github.com/juju/loggo

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o RelayPortal .
