FROM golang:1.17.2@sha256:04a6b03beb96e280c832ad0155c9f8dcedc9903ca0c2fc3a7fec68b77e10a455 AS builder

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
