FROM node:12-alpine as dev

ENV CI=true

VOLUME /src
WORKDIR /src

RUN yarn install

ENTRYPOINT ["yarn", "start"]
