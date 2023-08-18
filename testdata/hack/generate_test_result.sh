#!/bin/bash

set -eux

proto_dir=$1

cd $(dirname $0)

go build -o protoc-gen-vlossom ../../cmd/protoc-gen-vlossom/main.go

protoc \
  --plugin=protoc-gen-vlossom=./protoc-gen-vlossom \
  --vlossom_out=../cases/$proto_dir \
  -I ../../proto \
  -I ../cases/$proto_dir \
  ../cases/$proto_dir/*.proto

rm protoc-gen-vlossom
