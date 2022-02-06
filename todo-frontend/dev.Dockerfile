FROM node:14.3.0-alpine3.10

WORKDIR /usr/src/app

EXPOSE 8080

CMD ["npm","run","serve"]