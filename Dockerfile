FROM golang:1.15-alpine3.12 AS build

WORKDIR /build

RUN apk --no-cache add gcc g++ make git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app

FROM alpine:3.12

EXPOSE 8081

RUN apk --no-cache add ca-certificates

WORKDIR /run

COPY --from=build /build/app ./

ENTRYPOINT /run/app