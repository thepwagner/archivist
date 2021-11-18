FROM golang:1.17.3@sha256:fb5d73f7ffea9cf8607c95ec793efe7591daa8f54af6ce84776b3d8afda10545 AS builder

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
