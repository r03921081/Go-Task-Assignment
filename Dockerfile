ARG ALPINEVERSION=3.20

FROM golang:1.21-alpine${ALPINEVERSION} as builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o server .

FROM alpine:${ALPINEVERSION}
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/server .

CMD ["./server"]

EXPOSE 8080
