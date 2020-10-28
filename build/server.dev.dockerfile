FROM golang:1.15.0-alpine3.12

ENV CGO_ENABLED=0

VOLUME /src
WORKDIR /src

RUN apk update && \
    apk add git && \
    go get -u github.com/cosmtrek/air && \
    go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate && \
    go get -u github.com/go-bindata/go-bindata/...

COPY air.toml /etc/air.toml

ENTRYPOINT ["air", "-d", "-c", "/etc/air.toml"]
