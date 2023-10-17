#!/bin/bash

cd ./AB/webui
npm install && npm run build
mv ./dist ../../bin/dist
cd ../../go_backend
go build -o ../bin/GotoBangumi 