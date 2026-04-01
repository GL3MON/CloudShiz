FROM node:current-alpine3.23

WORKDIR /app

COPY . .
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npx", "nodemon", "server.js"]