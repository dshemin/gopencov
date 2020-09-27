#!/bin/sh
#
# Run `migrate` cli with required configuration.
# A wrapper for `https://github.com/golang-migrate/migrate` command.
#
# IMPORTANT: Make sure that you had run `make run` before.
#
DIR="$(cd $(dirname "$0");pwd)"

set -a
. "$DIR/vars.env"
. "$DIR/common.sh"
set +a

docker_compose exec migrate -path "/src/migrations/$DB_DRIVER" -database "$DB_URI" $@
