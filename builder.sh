#!/bin/bash

echo -e "[!] Bulding versions.. "
echo -e "[!] Build linux.. "
GOARCH=amd64 GOOS=linux go build $file -o xmlrpcscan_linux
echo "[!] Done linux"

echo -e "[!] Build Darwin.. "
GOARCH=amd64 GOOS=darwin go build $file -o xmlrpcscan_osx
echo "[!] Done Darwin"

echo -e "[!] Build Windows.. "
GOARCH=amd64 GOOS=windows go build $file -o xmlrpcscan_windows
echo "[!] Done Windows..."