FROM golang:1.17.6@sha256:204c2e4637f7bcd0ebb32886d6f2d03f2deda92cf2f77e5a76746a77ff8c1de8 AS builder

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
