FROM golang:1.18.0@sha256:b5a9ae5907066849cbb2b47af93713b1f030e3eb6b36d84f2ae5f354b25c80d4 AS builder

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
