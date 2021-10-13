FROM golang:1.17.2@sha256:2ca682b8e030b4bfb5ae887f984c258d6c5ceb0be73cede8d0048e16656af2b5 AS builder

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
