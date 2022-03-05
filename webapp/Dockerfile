FROM golang:1.17 AS builder
WORKDIR /go/src/github.com/trelore/formula1/webapp
COPY . /go/src/github.com/trelore/formula1/webapp
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/trelore/formula1/webapp/app ./app
CMD ["./app"]