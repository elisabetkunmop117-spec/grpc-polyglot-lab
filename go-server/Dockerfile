
FROM golang:1.25.9-alpine AS builder

WORKDIR /app

COPY go-server/go.mod go-server/go.sum ./go-server/

WORKDIR /app/go-server

RUN go mod download

COPY go-server/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /server ./server

EXPOSE 50051

CMD ["./server"]
