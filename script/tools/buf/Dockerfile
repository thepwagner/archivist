FROM alpine:3.15.0@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300 AS buf

# renovate: datasource=github-releases depName=bufbuild/buf versioning=semver
ARG BUF_VERSION=v1.0.0-rc9
ARG BUF_CHECKSUM=c103dba1b06ae972f5f9f7de6c792824f48b6d4840fb0a336d3924fd3577a4ef

ARG BUF_MINISIGN_KEY=RWQ/i9xseZwBVE7pEniCNjlNOeeyp4BQgdZDLQcAohxEAH5Uj5DEKjv6
ARG BUFF_URL=https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-Linux-x86_64
ARG BUFF_DIGESTS_URL=https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/sha256.txt
RUN apk --no-cache add curl \
  && apk --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community add minisign \
  && mkdir -p /buf \
  && curl -Lo /buf/sha256.txt "$BUFF_DIGESTS_URL" \
  && curl -Lo /buf/sha256.txt.minisig "$BUFF_DIGESTS_URL.minisig" \
  && minisign -Vm /buf/sha256.txt -P "$BUF_MINISIGN_KEY" \
  && grep -q "$BUF_CHECKSUM" /buf/sha256.txt \
  && curl -o /usr/local/bin/buf -L "$BUFF_URL" \
  && echo "$BUF_CHECKSUM  /usr/local/bin/buf" | sha256sum -c \
  && chmod +x /usr/local/bin/buf \
  && apk del curl minisign

FROM golang:1.17.6@sha256:1e1e3fb5d4a72e75953785bba1f48b72353ac8467dad3f3af2995e99099dd994 AS protoc-gen-go
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o /protoc-gen-go github.com/golang/protobuf/protoc-gen-go


FROM buf AS protoc
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
