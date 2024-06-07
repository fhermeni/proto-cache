#!/bin/sh
VERSION=v1
SRC_DIR=protos/v1
PKG=github.com/fhermeni/proto-cache/v1
DST_DIR=go/v1

set -x
protoc --go_out=${DST_DIR} --proto_path=${SRC_DIR} --go_opt=module=${PKG} ${SRC_DIR}/other/types.proto
echo $?
protoc --go_out=${DST_DIR} --proto_path=${SRC_DIR} --go_opt=module=${PKG} ${SRC_DIR}/my/my.proto
echo $?

