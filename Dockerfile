FROM golang:1.17.5@sha256:eb61213fe612b6af346cab13c2b81f1b8113f9ccf23a5ca6b54fede8ffd63bc7 AS builder

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
