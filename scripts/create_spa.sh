#!/bin/sh

cd .. && swag init -g main.go
cd frontend && npm i && npm run build
cd .. && go build