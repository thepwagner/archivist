FROM golang:1.17.6@sha256:d7f2f6f649920ec58af3ecd82c460c04c7eb5335dfed57a5383ba60d83b5d0a8 AS builder

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
