FROM golang:1.17.1@sha256:6bf3a48a203d230b5e4cab749d44a0d8d51905c023f1ebe29927cffa43061f2c AS builder

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
