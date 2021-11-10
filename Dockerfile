FROM golang:1.17.3@sha256:b5bfe0255e6fac7cec1abd091b5cc3a5c40e2ae4d09bafbe5e94cb705647f0fc AS builder

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
