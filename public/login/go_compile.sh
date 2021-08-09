#!/usr/bin/env bash

# $1 - 工程目录
# $2 - 工程名称

scriptFilePath=$(cd `dirname $0`; pwd)
echo "[INFO] begin compile package : $scriptFilePath"

cd $scriptFilePath

protoc -I=$1 -I=. --go_out=plugins=grpc,,paths=source_relative:. client.proto

protoc -I=$1 -I=. --go_out=plugins=grpc,paths=source_relative:. login_service.proto


