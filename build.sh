#!/usr/bin/env bash
PROJ_DIR=$(cd ../;pwd)
export GOPATH=$GOPATH:$PROJ_DIR
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kubectl-info-x86 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o kubectl-info-darwin main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o kubectl-info-win main.go
