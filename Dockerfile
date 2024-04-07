FROM golang:1.18-alpine3.17 as builder
WORKDIR app
COPY . .
RUN go build -o server.bin .

FROM alpine:3.17
WORKDIR app
COPY --from=builder /go/app/server.bin .
CMD ["./server.bin"]
