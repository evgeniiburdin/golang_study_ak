#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: gofmt.sh <filename>"
  exit 1
fi

FILENAME=$1

go fmt "$FILENAME"


