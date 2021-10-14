FROM golang:1.17.2@sha256:cf615c1499e8bc2cf00696ba234cddd47fdb8d9a3b37b7c35726e46ee4ae08cc AS builder

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
