FROM golang:1.18.1@sha256:4da03e1af84d99724cda668c0b9b3af628bf9e5be8fe31793a94bf6f9f383b21 AS builder

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
