FROM golang:1.21-alpine as build

COPY . /build

WORKDIR /build

RUN go build -o /serve cmd/serve.go

FROM alpine:latest

RUN adduser -u 1000 -D app app

USER app

COPY --from=build /serve /home/app/serve

EXPOSE 8080

ENTRYPOINT ["/home/app/serve"]