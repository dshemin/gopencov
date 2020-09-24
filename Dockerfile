FROM golang:1.15.0-alpine3.12 as builder

ENV CGO_ENABLED=0

WORKDIR /gopencov

RUN apk update && apk add git

COPY . /gopencov

RUN go build \
        -v \
        -mod vendor \
        -o /gopencov/server \
        ./cmd/server/main.go

FROM scratch

COPY --from=builder /gopencov/server /gopencov

ENTRYPOINT ["/gopencov"]
