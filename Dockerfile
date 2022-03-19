FROM golang:1.18.0@sha256:4dd5d732a1b039de87296d3ff52f921d48a679c3670bb7ec48f38fcc1b5fefdc AS builder

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
