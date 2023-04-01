# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /APIs

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-apis

EXPOSE 8080

CMD [ "/docker-apis" ]