#!/bin/bash

echo Debug started...

FILENAME=$1

dlv debug "$FILENAME"

go build -o "myprogram.exe" main.go

dlv exec myprogram.exe

echo Debug ended.