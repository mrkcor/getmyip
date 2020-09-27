#!/usr/bin/env bash

set -e

GOOS=linux GOARCH=amd64 go build -o build/getmyip getmyip.go
tar czfv build/getmyip-linux-amd64.tar.gz -C build getmyip
rm build/getmyip

GOOS=linux GOARCH=386 go build -o build/getmyip getmyip.go
tar czfv build/getmyip-linux-386.tar.gz -C build getmyip
rm build/getmyip

GOOS=linux GOARCH=arm go build -o build/getmyip getmyip.go
tar czfv build/getmyip-linux-arm.tar.gz -C build getmyip
rm build/getmyip

GOOS=windows GOARCH=amd64 go build -o build/getmyip.exe getmyip.go
zip -j build/getmyip-windows-64.zip build/getmyip.exe
rm build/getmyip.exe

GOOS=windows GOARCH=386 go build -o build/getmyip.exe getmyip.go
zip -j build/getmyip-windows-32.zip build/getmyip.exe
rm build/getmyip.exe

GOOS=windows GOARCH=arm go build -o build/getmyip.exe getmyip.go
zip -j build/getmyip-windows-arm.zip build/getmyip.exe
rm build/getmyip.exe