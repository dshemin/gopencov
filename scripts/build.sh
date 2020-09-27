#!/bin/sh
#
# Build all required containers
#
DIR="$(cd $(dirname "$0");pwd)"

docker build -t dshemin/gopencov-server:latest -f "$DIR/../build/server.prod.dockerfile" .
docker build -t dshemin/gopencov-frontend:latest -f "$DIR/../build/frontend.prod.dockerfile" web/
