#!/bin/bash

set -eux

cd $(dirname $0)

protoc \
--go_out=../pkg/protoptions \
--go_opt=paths=source_relative \
-I ./ \
./*.proto

