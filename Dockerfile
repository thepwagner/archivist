FROM golang:1.17.2@sha256:941f387984d023f357e6099e87bd799a91a17b8df3563cc067f487630e7cb1a0 AS builder

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
