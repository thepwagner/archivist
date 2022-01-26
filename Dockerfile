FROM golang:1.17.6@sha256:6cc9796e7cde1ab9aad195443e65b945deb883c4b5acec4c68a99a3c9f88db34 AS builder

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
