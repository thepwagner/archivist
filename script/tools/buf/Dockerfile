FROM alpine:3.14.0@sha256:adab3844f497ab9171f070d4cae4114b5aec565ac772e2f2579405b78be67c96 AS buf

# renovate: datasource=github-releases depName=bufbuild/buf versioning=semver
ARG BUF_VERSION=v0.43.2
ARG BUF_CHECKSUM=53933427cb5518830a3fcf11d3bd9b6e3fb36eae9086731484b8ab50630c9d09

ARG BUFF_URL=https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-Linux-x86_64
RUN apk --no-cache add --virtual .build curl \
  && curl -o /usr/local/bin/buf -L "$BUFF_URL" \
  && echo "$BUF_CHECKSUM  /usr/local/bin/buf" | sha256sum -c \
  && chmod +x /usr/local/bin/buf \
  && apk del .build


FROM golang:1.16.7@sha256:5cdc91c9e67e7b95ef5a1c9156af24aaccbb4e339799efd441fd7e961b3e2e75 AS protoc-gen-go
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o /protoc-gen-go github.com/golang/protobuf/protoc-gen-go


FROM buf AS protoc
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
