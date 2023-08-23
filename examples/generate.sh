#!/bin/bash

set -eux

cd $(dirname $0)

go build -o protoc-gen-venus ../cmd/protoc-gen-venus/main.go

protoc \
  --plugin=protoc-gen-venus=./protoc-gen-venus \
  --venus_out=./ \
  --venus_opt=expose_all=true \
  -I ../proto \
  -I ./ \
  ./example.proto

rm protoc-gen-venus
