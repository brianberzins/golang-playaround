FROM golang:1.18 AS builder
WORKDIR /go/src/golang-playgound
COPY ./ ./
RUN go build ./...

FROM scratch
COPY --from=builder /go/src/golang-playgound/golang-playaround ./golang-playaround
CMD ["/golang-playaround"]
