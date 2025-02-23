FROM golang:latest

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /fyp-server

EXPOSE 8080

CMD ["/fyp-server"]