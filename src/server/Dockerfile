FROM golang:1.19-alpine as dev

WORKDIR /go/src/app

COPY ./go.mod go.mod
COPY ./go.sum go.sum
RUN go mod download
ENV WELLREAD_ENV=development

CMD go run main.go

# prod
FROM dev as prod
ENV WELLREAD_ENV=production
COPY . .

RUN go build

CMD ./wellread

