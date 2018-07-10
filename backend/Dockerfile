FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/ComputePractice2018/AddressBook/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/addressbook github.com/ComputePractice2018/AddressBook/backend

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/addressbook /opt/addressbook

EXPOSE 8080

ENV CONNECTION "addressbook:SuperSecretPassword@tcp(db:3306)/addressbook"
ENTRYPOINT /opt/addressbook -connection ${CONNECTION}