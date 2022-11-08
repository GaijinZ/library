# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /books

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/*.go ./
COPY pkg/ ./pkg

RUN go build -o /books-go

EXPOSE 8080

CMD ["/books-go"]
