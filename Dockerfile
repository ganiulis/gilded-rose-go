FROM golang:1.21-alpine

WORKDIR /usr/src/app

COPY . ./

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN go mod tidy \
 && go mod download \
 && go mod verify

CMD ["/bin/sh"]

