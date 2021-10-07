# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

#Stops windows? error
ENV CGO_ENABLED=0

RUN go build -o /gin-mysql

FROM scratch AS bin

WORKDIR /

COPY --from=build /gin-mysql /gin-mysql

EXPOSE 8080

ENTRYPOINT [ "/gin-mysql" ]
