#!/bin/zsh

go run ./cmd/grpc-sandbox-server/main.go

#protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative proto/hello/v1/hello.proto