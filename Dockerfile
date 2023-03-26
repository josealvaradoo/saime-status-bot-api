FROM golang:latest

# Set up the working directory
RUN mkdir /app
WORKDIR /app

# Install Air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copy the project
COPY . .

# Install modules
RUN go mod download

# Build
RUN go build -o /main src/cmd/main.go

EXPOSE 3000

CMD ["../main"]