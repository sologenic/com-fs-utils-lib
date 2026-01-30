#!/bin/bash
set -e

PROTO_DIR=$(dirname $(dirname "$0"))
cd $PROTO_DIR

# Generate Go files
for proto in $(find . -name "*.proto" -not -path "./node_modules/*"); do
  protoc \
    --proto_path=. \
    --proto_path="$(go env GOPATH)/src" \
    --go_out=. \
    --go_opt=paths=source_relative \
    "$proto"
done

# Generate TypeScript files
rm -rf node_modules
npm i

for proto in $(find . -name "*.proto" -not -path "./node_modules/*"); do
  protoc \
    --plugin=protoc-gen-ts_proto="./node_modules/.bin/protoc-gen-ts_proto" \
    --ts_proto_out=. \
    --ts_proto_opt=esModuleInterop=true \
    --proto_path=. \
    --proto_path="$(go env GOPATH)/src" \
    "$proto"
done

# Build TypeScript
npm run build
git add build/

git add *.ts
rm -rf node_modules