FROM golang:1.17.4@sha256:34d4eb5e52a33b6c9d67eda847efd7e9f5dd3a491cd76167c84f36c141d0d703 AS builder

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
