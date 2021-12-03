FROM golang:1.17.3@sha256:cc7b8f047a157a565ace90eaa2d4f2ebaa73b3c00e7159ee87e6b5036a91e059 AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ARG CGO_ENABLED=0
RUN go build -o /archivist .

FROM scratch
COPY go.sum /go.sum
COPY --from=builder /archivist /archivist
ENTRYPOINT ["/archivist"]
