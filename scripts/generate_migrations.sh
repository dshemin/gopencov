#!/bin/sh

DBS="postgresql"

for db in $DBS; do
  echo "Generate bin data for $db migrations"
  go-bindata -pkg $db \
    -o internal/$db/migrations.go \
    -prefix ../../migrations/$db \
    ../../migrations/$db/...
done
