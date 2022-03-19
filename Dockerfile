FROM golang:1.18.0@sha256:975f5ba172bf536b32a30642f1cb1915a2cda58e94a15926cf7fae2219a5f5dd AS builder

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
