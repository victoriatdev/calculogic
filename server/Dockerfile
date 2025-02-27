FROM golang:alpine as build-env

WORKDIR /go/src
COPY . /go/src/fyp-server
RUN cd /go/src/fyp-server && go build .

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/fyp-server/fyp-server /app
COPY .env /app

EXPOSE 8080

CMD ["./fyp-server"]