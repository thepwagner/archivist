FROM golang:1.17.6@sha256:f71d4cabfdc092c07c7973cbdf8d6787f85e80a2f8ee78a7491949708215a4d3 AS builder

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
