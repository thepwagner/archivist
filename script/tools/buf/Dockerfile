FROM alpine:3.11 AS buf

ARG BUF_VERSION=v0.32.1
ARG BUF_CHECKSUM=8f4f3ff4b517f436f98e915812f00357724bca9bd601a36e6214584eb8b9d1ea

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
