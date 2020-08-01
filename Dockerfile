FROM golang:alpine AS build-env

WORKDIR /go/src/github.com/yebh17/go-app

COPY . .

RUN go build -o main

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*

WORKDIR /go/src/github.com/yebh17/go-app

COPY --from=build-env /go/src/github.com/yebh17/go-app* /go/src/github.com/yebh17/go-app

EXPOSE 6777

ENTRYPOINT [ "./main" ]
