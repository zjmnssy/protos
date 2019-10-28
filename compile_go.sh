#!/usr/bin/env bash

scriptFilePath=$(cd `dirname $0`; pwd)
echo "[INFO] current work path : $scriptFilePath"

PROJECT_FATHER_PATH=${scriptFilePath%/*}

function goModInit()
{
    isInited=`find -name *.mod`
    if [ -z $isInited ] 
    then 
        echo "[INFO] project need init"
        go mod init protos
    else
        echo "[INFO] project not need init"
    fi
}

function compileWithImport()
{
    protoFile=$1"/"$2
    protoc -I=$PROJECT_FATHER_PATH -I=$1 --go_out=$PROJECT_FATHER_PATH $protoFile
}

function compileWithoutImport()
{
    protoFile=$1"/"$2
    protoc -I=$1 --go_out=$1 $protoFile
}

function compileProto()
{
    protoFile=$1"/"$2
    echo  "[INFO] compile $protoFile."
    if grep -q "import" $protoFile
    then
        compileWithImport $1 $2
    else
        compileWithoutImport $1 $2
    fi
}

function compileMain()
{
    for file in `ls $1`
    do
        if [ -d $1"/"$file ] ; then
            compileMain $1"/"$file
        else
            if [[ "$file" =~  "proto"  ]] ; then
                compileProto $1 $file
            else 
                echo  "[WARN] $file not proto file, ignore."
            fi
        fi
    done
}

goModInit
compileMain $scriptFilePath


