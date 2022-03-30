FROM golang:1.18.0@sha256:60b6f0049fa1d6cbcc0b2062173ffc93e842d097e0e62974afff08fa6f1fe9e9 AS builder

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
