#!/bin/bash

set -eux

cd $(dirname $0)

go build -o protoc-gen-vlossom ../../cmd/protoc-gen-vlossom/main.go

proto_dirs=$(ls ../cases)

for proto_dir in $proto_dirs; do
  protoc \
    --plugin=protoc-gen-vlossom=./protoc-gen-vlossom \
    --vlossom_out=../cases/$proto_dir \
    --vlossom_opt=pretty_output=true \
    -I ../../proto \
    -I ../cases/$proto_dir \
    ../cases/$proto_dir/entry.proto
done

rm protoc-gen-vlossom
