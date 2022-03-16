FROM golang:1.18.0@sha256:7a3cc1bd39e3937b4eddd2e66c4ed7c1852a4e7fa42b73507ed4ee50c02978a4 AS builder

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
