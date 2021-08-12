FROM golang:1.16.7@sha256:3c4de86eec9cbc619cdd72424abd88326ffcf5d813a8338a7743c55e5898734f AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ARG CGO_ENABLED=0
RUN go build -o /archivist .

FROM alpine:3.14.1@sha256:eb3e4e175ba6d212ba1d6e04fc0782916c08e1c9d7b45892e9796141b1d379ae
COPY --from=builder /archivist /usr/local/bin
ENTRYPOINT ["/usr/local/bin/archivist"]
