#!/bin/sh
set -eu

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BACKEND_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
PROTO_DIR="$(cd "${BACKEND_DIR}/../steam_protobufs" && pwd)"

cd "${BACKEND_DIR}"

echo "Generating Go types from protobufs..."
protoc \
	--go_out=steamapi \
	--go_opt=paths=source_relative \
	--go_opt=Mcommon_base.proto=github.com/Alia5/steaminputdb.com/steamapi \
	--go_opt=Mcommon.proto=github.com/Alia5/steaminputdb.com/steamapi \
	--go_opt=Menums.proto=github.com/Alia5/steaminputdb.com/steamapi \
	--go_opt=Mservice_steaminputmanager.proto=github.com/Alia5/steaminputdb.com/steamapi \
	--go_opt=Mservice_publishedfile.proto=github.com/Alia5/steaminputdb.com/steamapi \
	--go_opt=Mservice_storequery.proto=github.com/Alia5/steaminputdb.com/steamapi \
	--proto_path="${PROTO_DIR}/webui" \
	--proto_path="${PROTO_DIR}/steam" \
	"${PROTO_DIR}/steam/enums.proto" \
	"${PROTO_DIR}/webui/common_base.proto" \
	"${PROTO_DIR}/webui/common.proto" \
	"${PROTO_DIR}/webui/service_publishedfile.proto" \
	"${PROTO_DIR}/webui/service_storequery.proto"
