FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/AddressBook/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN apk add --no-cache git
RUN go get -u github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/addressbook github.com/ComputePractice2018/AddressBook/backend

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/addressbook /opt/addressbook

EXPOSE 8080

ENV NAME "Ivan"
ENTRYPOINT /opt/addressbook -name ${NAME}