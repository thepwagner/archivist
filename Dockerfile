FROM golang:1.18.1@sha256:bf8b2a996913b16e1b47f93d59694188c64301e02624e9ad8f65c54d97b0c5c5 AS builder

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
