FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/server .

CMD ["./server"]

EXPOSE 8080
