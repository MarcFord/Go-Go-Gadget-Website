#!/bin/bash
echo "Building the web app now"
cd web
[ ! -e web ] || rm web
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web .
if [ $? -eq 0 ]
then
    echo "Web App built to start the server run 'make start_srv'"
else
    echo "There was a problem building the web app!"
fi