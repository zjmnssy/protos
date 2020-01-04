#!/usr/bin/env bash

scriptFilePath=$(cd `dirname $0`; pwd)
echo "[INFO] begin compile package : $scriptFilePath"

cd $scriptFilePath

protoc -I=. --go_out=plugins=grpc:. error_code.proto


