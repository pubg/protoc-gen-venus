#!/bin/bash

set -eux

cd $(dirname $0)

protoc \
--go_out=../generator/ \
--go_opt=paths=source_relative \
-I ./ \
./options.proto

