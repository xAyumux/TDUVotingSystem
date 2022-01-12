FROM node:16.13.1

WORKDIR /app
COPY package*.json ./

RUN yarn global add @vue/cli-service && yarn install

COPY /client /app/

