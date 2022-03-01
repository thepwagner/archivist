FROM golang:1.17.7@sha256:c2ca4729eae792e8066e471e4b71d61a69b22ad71ca31e9d524fde5addabbd7e AS builder

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
