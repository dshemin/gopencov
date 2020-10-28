#!/bin/sh
#
# Run development environment.
# This command will run all required parts of the application with live reloading
# and all third-parties services.
#
DIR="$(cd $(dirname "$0");pwd)"

set -a
. "$DIR/vars.env"
. "$DIR/common.sh"
set +a

docker_compose up --build
