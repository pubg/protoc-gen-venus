#!/bin/bash

set -eux

cd $(dirname $0)

protoc \
--plugin=protoc-gen-vlossom=prodoc-gen-debug \
--debug_out=./ \
--debug_opt=dump_file=request.pb.bin \
--debug_opt=parameter=expose_all=true \
-I ../proto \
-I ./ \
./example.proto
