FROM golang:1.16.2 AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ARG CGO_ENABLED=0
RUN go build -o /archivist .

FROM alpine:3.11
COPY --from=builder /archivist /usr/local/bin
ENTRYPOINT ["/usr/local/bin/archivist"]
