#!/bin/zsh

rm -i -rf gen/proto

protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative proto/v1/*.proto