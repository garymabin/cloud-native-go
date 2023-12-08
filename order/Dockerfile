FROM golang:1.15.0 as golang-builder
MAINTAINER binma@thoughtworks.com

ENV SOURCES /go/src/github.com/garymabin/cloud-native-go/order

COPY . ${SOURCES}

RUN cd ${SOURCES} && go get github.com/gorilla/mux

RUN cd ${SOURCES} && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=golang-builder /go/src/github.com/garymabin/cloud-native-go/order/app .

ENV PORT 8080

ENV PAYMENT_SERVICE localhost

EXPOSE 8080

ENTRYPOINT /root/app
