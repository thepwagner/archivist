FROM golang:1.17.5@sha256:0502dc53f72f2b6fd3c3a6d4ee39932355c98da06ab18e9a20dec32fcb2ff994 AS builder

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
