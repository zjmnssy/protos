#!/usr/bin/env bash

scriptFilePath=$(cd `dirname $0`; pwd)
echo "[INFO] current work path : $scriptFilePath"

command=$1

function cleanGoModule()
{
    echo "                                      clean go module"
    rm $scriptFilePath/go.mod
    rm $scriptFilePath/go.sum
    echo "       -------------------------------------------------------------------------------        "
}

function cleanGo()
{
    for file in `ls $1`
    do
        if [ -d $1"/"$file ] ; then
            cleanGo $1"/"$file
        else
            if [[ "$file" =~  ".pb.go"  ]] ; then
                protoFile=$1"/"$file
                echo "rm $protoFile"
                rm $protoFile
            fi
        fi
    done
}

function cleanCpp()
{
    for file in `ls $1`
    do
        if [ -d $1"/"$file ] ; then
            cleanCpp $1"/"$file
        else
            if [[ "$file" =~  ".cc"  ]] ; then
                protoFile=$1"/"$file
                echo "rm $protoFile"
                rm $protoFile
            fi

            if [[ "$file" =~  ".h"  ]] ; then
                protoFile=$1"/"$file
                echo "rm $protoFile"
                rm $protoFile
            fi
        fi
    done
}

function cleanMain() 
{
    echo "$1 11111111111"
    if [[ "$1" =~  "cpp"  ]] ; then
        echo "                                     clean compile proto cpp files"
        cleanCpp $scriptFilePath
    elif [[ "$1" =~  "go"  ]] ; then
        cleanGoModule
        echo "                                     clean compile proto go files"
        cleanGo $scriptFilePath
    else 
        echo "[WARN] unknow type of $1"
    fi
}

echo "**********************************************************************************************"
cleanMain $command
echo "**********************************************************************************************"
