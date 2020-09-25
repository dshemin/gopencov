FROM node:12-alpine as builder

WORKDIR /frontend

COPY . /frontend

RUN apk add --no-cache \
    && apk add --no-cache \
    && yarn install \
    && yarn build

FROM nginx:1.19.2-alpine

COPY --from=builder /frontend/build /usr/share/nginx/html
