#!/bin/bash
set -e

mkdir -p src/lib/steamapi/proto

protoc \
    --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=src/lib/steamapi/proto \
    --ts_proto_opt=esModuleInterop=true \
    --ts_proto_opt=fileSuffix=.pb \
    --proto_path=../steam_protobufs/webui \
    --proto_path=../steam_protobufs/steam \
    ../steam_protobufs/webui/common_base.proto \
    ../steam_protobufs/webui/common.proto \
    ../steam_protobufs/webui/service_publishedfile.proto \
    ../steam_protobufs/webui/service_storequery.proto
