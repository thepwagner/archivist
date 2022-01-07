FROM golang:1.17.6@sha256:d36ec9839e6ebac63a53f3e15758ffa339b81dc5df6c9d41a18a3f9302bd0d90 AS builder

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
