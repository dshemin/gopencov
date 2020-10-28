#!/bin/sh
#
# Run `psql` to the development database.
#
# IMPORTANT: Make sure that you had run `make run` before.
#
DIR="$(cd $(dirname "$0");pwd)"

set -a
. "$DIR/vars.env"
. "$DIR/common.sh"
set +a

docker_compose exec postgres psql "$DB_URI" $@
