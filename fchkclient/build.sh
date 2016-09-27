#!/bin/sh
if [ "$1" = "w" ]; then 
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v 
elif [ "$1" = "m" ]; then
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v 
else
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v
fi

