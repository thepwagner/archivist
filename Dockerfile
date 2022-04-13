FROM golang:1.18.1@sha256:91ee8a9baa730812ae25daf520ecb29750c979e10e8454e6ccbe478a4c8fbb22 AS builder

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
