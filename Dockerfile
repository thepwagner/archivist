FROM golang:1.17.3@sha256:80840b4af07c63588dc00be0d4e99dfacc6453f54b1dbc103c35e3033ec40dcd AS builder

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
