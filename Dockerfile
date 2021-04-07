FROM golang:1.6.3 AS builder
WORKDIR /go/src/github.com/alicerum/gotest/
COPY . .
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/alicerum/gotest/gotest .
CMD ["./gotest"]
