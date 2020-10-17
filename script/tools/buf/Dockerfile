FROM alpine:3.11 AS buf

ARG BUF_VERSION=v0.27.1
ARG BUF_CHECKSUM=2940c9826448c143d62c66264d4e111e4f74a0d9292b15ba919916fe3827a3fc

ARG BUFF_URL=https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-Linux-x86_64
RUN apk --no-cache add --virtual .build curl \
  && curl -o /usr/local/bin/buf -L "$BUFF_URL" \
  && echo "$BUF_CHECKSUM  /usr/local/bin/buf" | sha256sum -c \
  && chmod +x /usr/local/bin/buf \
  && apk del .build


FROM golang:1.15.2 AS protoc-gen-go
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o /protoc-gen-go github.com/golang/protobuf/protoc-gen-go


FROM buf AS protoc
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
