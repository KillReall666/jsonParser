FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
COPY ports.json ./
COPY cmd ./cmd
COPY internal ./internal
COPY docs ./docs

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Run
CMD ["/app/main"]