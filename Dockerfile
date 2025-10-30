FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

# FROM alpine:3.19

# WORKDIR /app

# COPY --from=builder /app .

# RUN chmod +x /app/main
# EXPOSE 3000

# CMD ["go run app/cmd/main.go"]
