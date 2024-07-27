FROM node:22.5-alpine as node

WORKDIR /home/node/app

COPY /package*.json ./

RUN npm i

COPY . ./

EXPOSE 4001

CMD ["node", "./src/index.js"]
