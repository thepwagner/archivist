FROM alpine:3.13.4 AS buf

ARG BUF_VERSION=v0.41.0
ARG BUF_CHECKSUM=a8b597394160d00499f5c12326e64e7456090fb5869d97ccd816528b74537956

ARG BUFF_URL=https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-Linux-x86_64
RUN apk --no-cache add --virtual .build curl \
  && curl -o /usr/local/bin/buf -L "$BUFF_URL" \
  && echo "$BUF_CHECKSUM  /usr/local/bin/buf" | sha256sum -c \
  && chmod +x /usr/local/bin/buf \
  && apk del .build


FROM golang:1.16.3 AS protoc-gen-go
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o /protoc-gen-go github.com/golang/protobuf/protoc-gen-go


FROM buf AS protoc
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
