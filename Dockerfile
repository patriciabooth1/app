FROM golang:1.16.3-alpine

RUN apk add --no-cache make

WORKDIR /src/app

COPY . .

RUN make build

EXPOSE 8123

CMD ["/src/app/./executable"]