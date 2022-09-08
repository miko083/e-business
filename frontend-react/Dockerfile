FROM node:18.8.0

WORKDIR /ebiznes-consoleshop-frontend

COPY . .

RUN npm ci
RUN npm run build

EXPOSE 3000

CMD [ "npx", "serve", "build" ]