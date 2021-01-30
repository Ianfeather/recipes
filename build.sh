#!/usr/bin/env bash

rm -rf build/*
go fmt ./...
go test ./... -v
GOOS=linux GOARCH=amd64 go build -o main cmd/main.go
zip build/main.zip main
rm -rf main
