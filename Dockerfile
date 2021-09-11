FROM golang:1.16-alpine as base

# setup directories
RUN mkdir /app
ADD . /app
WORKDIR /app

## pull dependencies
RUN go mod download

## build app
RUN go build -o main .

## start application
CMD ["/app/main"]
