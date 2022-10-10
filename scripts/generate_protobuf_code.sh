#!/bin/bash

protoc -I=./classificator \
  --python_out=classificator/src/etl/proto \
  --python_out=classificator/src/email_classificator/proto \
  classificator/*.proto

protoc -I=./classificator \
  --go_out=classificator/src/email_filter/proto \
  --go_out=classificator/src/stats_storer/proto \
  --go_out=classificator/src/exitpoint/proto \
  --go_opt=paths=source_relative classificator/*.proto \

echo "Done"
