#!/bin/bash

# shellcheck disable=SC2086

## USAGE:
#    ./build_krt.sh <new-version-name>

set -eu

VERSION_DIR="classificator"

# NOTE: if yq commands fails it due to the awesome Snap installation that is confined (heavily restricted).
# Please install yq binary from https://github.com/mikefarah/yq/releases and think twice before using Snap next time.

echo -e "Reading current version: \c"
CURRENT_VERSION=$(yq e .version ${VERSION_DIR}/krt.yml)

echo "${CURRENT_VERSION}"

VERSION=${VERSION_DIR}-${1:-${CURRENT_VERSION#$VERSION_DIR-}}

if [ -z "$VERSION" ]; then
  echo "error setting KRT version"
  exit 1;
fi

echo "Building Stats Storer node Golang binary..."
cd classificator/src/stats_storer
go build -o ../../bin/stats_storer .

echo "Building Repairs Handler node Golang binary..."
cd ../repairs_handler
go build -o ../../bin/repairs_handler .

echo "Building Exitpoint node Golang binary..."
cd ../exitpoint
go build -o ../../bin/exitpoint .
cd ../../..


echo "Generating $VERSION.krt..."

mkdir -p build/${VERSION_DIR}
rm -rf ./build/${VERSION_DIR}/{docs,src,assets,models,*.proto,*.yml} || true

cd build/${VERSION_DIR}

cp  -r ../../${VERSION_DIR}/* .

yq eval --inplace -- ".version = \"${VERSION}\"" ./krt.yml

tar -zcf ../${VERSION}.krt  --exclude=*.krt --exclude=*.tar.gz *
cd ../../
rm -rf build/${VERSION_DIR} || true

echo "Done"
