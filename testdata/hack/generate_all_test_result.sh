#!/bin/bash

set -eux

cd $(dirname $0)

go build -o protoc-gen-venus ../../cmd/protoc-gen-venus/main.go

proto_dirs=$(ls ../cases)

for proto_dir in $proto_dirs; do
  protoc \
    --plugin=protoc-gen-venus=./protoc-gen-venus \
    --venus_out=../cases/$proto_dir \
    --venus_opt=pretty_output=true \
    -I ../../proto \
    -I ../cases/$proto_dir \
    ../cases/$proto_dir/entry.proto
done

rm protoc-gen-venus
