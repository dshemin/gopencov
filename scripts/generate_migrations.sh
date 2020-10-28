#!/bin/sh
#
# Use go-bindata for packing all database migration into the application.
#

DBS="postgres"

for db in $DBS; do
  echo "Generate bin data for $db migrations"
  go-bindata -pkg $db \
    -o internal/$db/migrations.go \
    -prefix ../../migrations/$db \
    ../../migrations/$db/...
done
