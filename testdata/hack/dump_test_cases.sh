#!/bin/bash

set -eux

cd $(dirname $0)

proto_dirs=$(ls ../cases)

for proto_dir in $proto_dirs; do
  protoc \
    --plugin=protoc-gen-vlossom=prodoc-gen-debug \
    --debug_out=../cases/$proto_dir \
    --debug_opt=dump_binary=true \
    -I ../../proto \
    -I ../cases/$proto_dir \
    ../cases/$proto_dir/entry.proto
done

