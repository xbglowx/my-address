FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v . \
    && CGO_ENABLED=0 GOOS=linux go build -o app .

FROM scratch
COPY --from=builder /go/src/app/app /
ENTRYPOINT ["/app"]
