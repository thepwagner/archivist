FROM golang:1.17.3@sha256:dac1fd50dc298852005ed2d84baa4f15ca86b1042fce6d8da3f98d6074294bf4 AS builder

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
