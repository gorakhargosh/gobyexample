#!/bin/sh

TOP_DIR=$(git rev-parse --show-toplevel)

GO=${TOP_DIR}/bin/go

export GOPATH=${TOP_DIR}/third_party/golang:$GOPATH


$GO get github.com/gin-gonic/gin
