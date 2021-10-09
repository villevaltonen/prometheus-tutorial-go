FROM golang:1.16-alpine as base

# setup directories
RUN mkdir /app
ADD . /app
WORKDIR /app

## pull dependencies
RUN go mod download

## build app
RUN go build -o main .

## create the runtime container
FROM alpine:3.14

## copy the binary from the base image
COPY --from=base /app/main /bin/main

## use a non-root user
RUN adduser --disabled-password -u 1000 user
USER user

ENTRYPOINT [ "/bin/main" ]
