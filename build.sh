#!/bin/sh
SRC_DIR=src
PKG=github.com/fhermeni/proto-cache
DST_DIR=./
protoc --go_out=${DST_DIR} --proto_path=${SRC_DIR} --go_opt=module=${PKG} ${SRC_DIR}/other/types.proto
echo $?
protoc --go_out=${DST_DIR} --proto_path=${SRC_DIR} --go_opt=module=${PKG} ${SRC_DIR}/my/my.proto
echo $?

