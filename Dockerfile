FROM golang:1.17.6@sha256:fff998d0f7a46baeb15d8478ee909a8237377e07ccf1a8e77d77a32033021d62 AS builder

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
