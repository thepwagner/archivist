FROM golang:1.17.0@sha256:50b5f2673c0995eb3cda07150611acd3bd0523b0b3da5bd079e9b9d905f8e1cc AS builder

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
