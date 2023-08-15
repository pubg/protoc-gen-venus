#!/bin/bash

set -eux

cd $(dirname $0)

proto_dirs=$(ls ../cases)

for proto_dir in $proto_dirs; do
  protoc \
    --plugin=protoc-gen-vlossom=./redirect \
    --vlossom_out=./ \
    -I ../../proto \
    -I ../cases \
    ../cases/$proto_dir/*.proto
  mv request.pb.bin ../cases/$proto_dir/request.pb.bin
done

