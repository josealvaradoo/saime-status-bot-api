FROM golang:latest AS builder

# Install Air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copy the project
RUN mkdir /app
COPY . /app
WORKDIR /app

# Install modules
RUN go mod download

EXPOSE 3000

CMD [ "air"]