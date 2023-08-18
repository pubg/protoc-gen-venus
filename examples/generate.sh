#!/bin/bash

set -eux

cd $(dirname $0)

go build -o protoc-gen-vlossom ../cmd/protoc-gen-vlossom/main.go

protoc \
  --plugin=protoc-gen-vlossom=./protoc-gen-vlossom \
  --vlossom_out=./ \
  --vlossom_opt=expose_all=true \
  -I ../proto \
  -I ./ \
  ./example.proto

rm protoc-gen-vlossom
