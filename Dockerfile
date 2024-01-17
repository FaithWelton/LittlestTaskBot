FROM --platform=linux/arm golang:1.21 as builder
WORKDIR /app
ADD . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app main.go

FROM alpine:3.18 as production
COPY --from=builder /app .

CMD ["./app"]