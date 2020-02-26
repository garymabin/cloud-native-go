FROM golang:1.12.7-alpine as golang-builder
MAINTAINER binma@thoughtworks.com

ENV SOURCES /go/src/github.com/garymabin/cloud-native-go

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=golang-builder /go/src/github.com/garymabin/cloud-native-go/app .

ENV PORT 8080

EXPOSE 8080

ENTRYPOINT /root/app
