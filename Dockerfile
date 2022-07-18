FROM golang:1.18 AS builder
WORKDIR /go/src/golang-playaround
ENV CGO_ENABLED=0 GOOS=linux
COPY ./ ./
RUN go build ./...

FROM scratch
COPY --from=builder /go/src/golang-playaround/golang-playaround /
CMD ["/golang-playaround"]
