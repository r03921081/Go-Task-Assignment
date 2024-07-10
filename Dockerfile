ARG ALPINEVERSION=3.20

FROM golang:1.21-alpine${ALPINEVERSION} as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server

FROM alpine:${ALPINEVERSION}
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

CMD ["./server"]
