#!/usr/bin/env bash

scriptFilePath=$(cd `dirname $0`; pwd)
echo "[INFO] current work path : $scriptFilePath"

PROJECT_FATHER_PATH=${scriptFilePath%/*}
PROJECT_PATH=$scriptFilePath
PACKAGE_NAME="github.com/zjmnssy/protos"

function goModInit(){
    isInited=`find -name *.mod`
    if [ -z $isInited ] 
    then 
        echo "[INFO] project need init"
        go mod init $PACKAGE_NAME
        echo "" >> go.mod
        echo "replace $PACKAGE_NAME => $scriptFilePath" >> go.mod
    else
        echo "[INFO] project not need init"
    fi
}

function compileMain()
{
    for file in `ls $1`
    do
        if [ -d $1"/"$file ] ; then
            compileMain $1"/"$file
        else
            if [[ "$file" =~  ".proto"  ]] ; then
                fileFull=$1"/"$file
                replacement="."
                stripPrefix=${fileFull/#$scriptFilePath/$replacement}
                echo "now compile $stripPrefix"

                protoc  --go_out=plugins=grpc,paths=source_relative:. $stripPrefix
            fi
        fi
    done
}

echo "**********************************************************************************************"
echo "                                    1. go module init"
goModInit
echo "       -------------------------------------------------------------------------------        "
echo "                                    2. compile proto files"
compileMain $PROJECT_PATH
echo "**********************************************************************************************"


