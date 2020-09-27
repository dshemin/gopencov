#!/bin/sh
#
# Contains some common parts for over scripts.
#

docker_compose() {
  echo "Run docker-compose $@"
  docker-compose \
    --file "$DIR/../deployments/docker-compose.dev.yml" \
    --env-file "$DIR/vars.env" \
    $@
}
