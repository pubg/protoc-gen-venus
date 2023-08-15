#!/bin/bash

set -eux

cd $(dirname $0)

protoc \
--plugin=protoc-gen-vlossom=./redirect \
--vlossom_out=./ \
--vlossom_opt=expose_all=true \
-I ../proto \
-I ./ \
./example.proto
