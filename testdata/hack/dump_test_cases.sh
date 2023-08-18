#!/bin/bash

set -eux

cd $(dirname $0)

proto_dirs=$(ls ../cases)

for proto_dir in $proto_dirs; do
  protoc \
    --plugin=protoc-gen-vlossom=prodoc-gen-debug \
    --debug_out=../cases/$proto_dir \
    --debug_opt=dump_binary=true \
    --debug_opt=dump_json=true \
    -I ../../proto \
    -I ../cases \
    ../cases/$proto_dir/*.proto
done

