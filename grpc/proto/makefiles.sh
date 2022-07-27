#!/bin/bash

PATH="$PATH:$(go env GOPATH)/bin"
export PATH

cd "$(dirname "$0")" || (echo "进入当前工作目录失败" && exit)


protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hello/hello.proto