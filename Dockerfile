FROM golang:1.19 as builder
WORKDIR /app
ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.18 as production
COPY --from=builder /app .

CMD ["./app"]