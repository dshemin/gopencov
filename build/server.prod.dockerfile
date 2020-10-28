FROM golang:1.15.0-alpine3.12 as builder

ENV CGO_ENABLED=0

RUN apk update && apk add git && go get -u github.com/go-bindata/go-bindata/...

COPY . /gopencov
WORKDIR /gopencov

RUN go generate ./... && \
    go build \
        -v \
        -mod vendor \
        -o /gopencov/server \
        ./cmd/server/main.go

FROM scratch

COPY --from=builder /gopencov/server /gopencov

ENTRYPOINT ["/gopencov"]
