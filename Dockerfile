FROM golang:alpine AS build
WORKDIR /go/src/SuperStar
COPY . .
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o superstar .

EXPOSE 3000

ENTRYPOINT ["/superstar"]