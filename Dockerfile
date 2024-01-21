FROM golang:1.21 as golang

RUN mkdir -p /

WORKDIR /

COPY . .

RUN make build