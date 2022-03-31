FROM golang:1.18.0@sha256:d3585bb87ea19d2260c53fcbad6cf0e6338be4d6d52e893e0cc055aa9a87d628 AS builder

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
