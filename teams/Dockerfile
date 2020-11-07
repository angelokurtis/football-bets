FROM node:lts-stretch-slim

ENV DATA_DIR=/usr/src/data
ADD ./data.tar.xz $DATA_DIR/

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install
COPY . .

EXPOSE 3000
CMD [ "node", "index.js" ]
