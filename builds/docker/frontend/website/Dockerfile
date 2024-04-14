FROM node:20.11.1-alpine3.19

WORKDIR /usr/src/website

COPY ./frontend/website/package.json .

RUN npm install

COPY ./frontend/website .

RUN npm run build
CMD [ "npm", "run" , "start" ]


