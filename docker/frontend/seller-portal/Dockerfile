FROM node:20.11.1-alpine3.19

WORKDIR /usr/src/seller-portal

COPY ./frontend/seller-portal/package.json .

RUN npm install

COPY ./frontend/seller-portal .

RUN npm run build
CMD [ "npm", "run" , "start" ]

