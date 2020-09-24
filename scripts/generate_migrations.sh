#!/bin/sh

echo "Generate bin data for with the PostgreSQL migrations"
go-bindata -pkg postgresql \
  -o internal/postgresql/migrations.go \
  -prefix ../../data/migrations/postgresql \
  ../../data/migrations/postgresql/...
