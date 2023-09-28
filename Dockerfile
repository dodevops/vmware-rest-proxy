FROM golang:1.21-alpine as build

COPY . /build

WORKDIR /build

RUN go build -o /serve cmd/serve.go

FROM alpine:latest

COPY --from=build /serve /

EXPOSE 8080

ENTRYPOINT ["/serve"]