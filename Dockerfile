FROM golang:1.18.1@sha256:102f43050460b26d3a1646169217813b30151373a843e603fb37a3f6998b8281 AS builder

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
