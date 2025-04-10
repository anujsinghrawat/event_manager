FROM golang:alpine

# Install git
RUN apk update && apk add --no-cache git

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy
