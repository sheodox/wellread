FROM golang:1.17-alpine as dev

WORKDIR /go/src/app

COPY ./src/server/go.mod go.mod
COPY ./src/server/go.sum go.sum
RUN go mod download

CMD go run main.go

# frontend build, used for prod only.
# in dev vite is run on host machine.
FROM node:16-alpine as frontend
WORKDIR /usr/src/frontend
COPY ./src/static/package*.json ./
RUN npm ci
ENV NODE_ENV=production
COPY ./src/static .
RUN npm run build

# prod
FROM dev as prod
WORKDIR /go/src/app
COPY --from=frontend /usr/src/frontend/dist /usr/src/frontend
COPY ./src/server/ ./

RUN go build

CMD ./wellread

