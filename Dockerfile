FROM golang:latest AS builder

# Set up the working directory
WORKDIR /app

# Install dependencies
COPY . .
RUN go mod download

# Build
RUN GOOS=linux GOARCH=amd64 go build -o /main src/cmd/main.go

# Run the application
FROM alpine

COPY --from=builder app/.env .
COPY --from=builder /main .

EXPOSE 3000

ENTRYPOINT ["/main"]