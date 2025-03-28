# build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o user-admin-api ./cmd/main.go

# run the binary
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-admin-api .

EXPOSE 8080

CMD ["./user-admin-api"]
