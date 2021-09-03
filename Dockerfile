FROM golang:1.17.0@sha256:c994ea4c0e524ea97ea7b4b21c19b968170a0c804b2fa7eee3c70c779fe84211 AS builder

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
