FROM golang:alpine AS build
WORKDIR /go/src/SuperStar
COPY . .
RUN go mod download