FROM golang:1.15.2 AS protoc-gen-go

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go build -o /protoc-gen-go github.com/golang/protobuf/protoc-gen-go


FROM alpine:3.11

ARG BUF_VERSION=v0.24.0
ARG BUF_CHECKSUM=c5a6e7784b5763d66f1b296b2579b21e5df3261a3865a6d9ccafbeaa42b77562

ARG BUFF_URL=https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-Linux-x86_64
RUN apk --no-cache add --virtual .build curl \
  && curl -o /usr/local/bin/buf -L "$BUFF_URL" \
  && echo "$BUF_CHECKSUM  /usr/local/bin/buf" | sha256sum -c \
  && chmod +x /usr/local/bin/buf \
  && apk del .build

COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
