FROM golang:alpine

WORKDIR /usr/src/cvapi

COPY ./code/go.sum ./
COPY ./code/go.mod ./

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN go mod download && go mod verify

COPY ./code .

RUN CGO_ENABLED=1 go build -v -o /usr/local/bin/cvapi ./...

ENTRYPOINT ["cvapi"]