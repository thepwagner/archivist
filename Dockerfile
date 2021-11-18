FROM golang:1.17.3@sha256:b29b38c8ccc4d755873e4bfb9cdc0e9bfcbdd48d62976fc80cf7082ec859d901 AS builder

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
