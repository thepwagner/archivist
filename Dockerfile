FROM golang:1.17.3@sha256:4eab694fe28880afaa10c436f8b726224580149e07c4f9c4cfeaa618b8ee0736 AS builder

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
