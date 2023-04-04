#!/bin/bash

VERSION_PATH="$PWD/classificator/src"
PY_CLASSIFICATOR_PATH="$VERSION_PATH/py-classificator"
GO_PY_CLASSIFICATOR_PATH="$VERSION_PATH/go-classificator"
COMMON_PATH="$VERSION_PATH/common"

protoc -I=./classificator \
  --python_out="$PY_CLASSIFICATOR_PATH/etl/proto" \
  --mypy_out="$PY_CLASSIFICATOR_PATH/etl/proto" \
  --python_out="$PY_CLASSIFICATOR_PATH/email_classificator/proto" \
  --mypy_out="$PY_CLASSIFICATOR_PATH/email_classificator/proto" \
  classificator/*.proto

protoc -I=./classificator \
  --go_out="$GO_PY_CLASSIFICATOR_PATH/etl/proto" \
  --go_out="$GO_PY_CLASSIFICATOR_PATH/email_classificator/proto" \
  --go_out="$COMMON_PATH/repairs_handler/proto" \
  --go_out="$COMMON_PATH/stats_storer/proto" \
  --go_out="$COMMON_PATH/exitpoint/proto" \
  --go_opt=paths=source_relative classificator/*.proto \

echo "Done"
