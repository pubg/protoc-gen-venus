#!/bin/bash

set -eux

cd $(dirname $0)

go build -o protoc-gen-venus ../main.go

export DEBUG=true

protoc \
  --plugin=protoc-gen-venus=./protoc-gen-venus \
  --venus_out=./ \
  -I ./ \
  -I ../proto \
  ./example.proto

rm protoc-gen-venus
