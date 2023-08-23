#!/bin/bash

set -eux

cd $(dirname $0)

protoc \
--plugin=protoc-gen-venus=prodoc-gen-debug \
--debug_out=./ \
--debug_opt=parameter=expose_all=true \
-I ../proto \
-I ./ \
./example.proto
