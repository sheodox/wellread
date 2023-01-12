FROM node:18-alpine as frontend
WORKDIR /usr/src/frontend
COPY ./package*.json ./
RUN npm ci
ENV NODE_ENV=production
COPY . .
RUN npm run build

CMD node build
