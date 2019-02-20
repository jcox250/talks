FROM golang:1.11.3-alpine3.8

EXPOSE 3999

WORKDIR /go/src/talks

COPY . .

RUN apk update && apk upgrade && apk add --no-cache \
        bash \
        git  

RUN go get golang.org/x/tools/cmd/present
RUN go install golang.org/x/tools/cmd/present

