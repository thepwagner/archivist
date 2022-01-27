FROM golang:1.17.6@sha256:9718826ce9dd3f16c54b9cc8467c9defd7148683dc13e919fb44d4883a25b9eb AS builder

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
