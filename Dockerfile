FROM golang:1.17.6@sha256:2adfd0d7f507cfe0dab9aefaf3e0b973b9c8cce48caa43431dca7f8a2cef5557 AS builder

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
