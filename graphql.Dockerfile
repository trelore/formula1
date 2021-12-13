FROM golang:1.17 AS builder
WORKDIR /go/src/github.com/trelore/formula1/
COPY ../ /go/src/github.com/trelore/formula1
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./formulagraphql

FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/trelore/formula1/app ./app
CMD ["./app"]