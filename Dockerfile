# Build stage
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod .
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o web-app .

# Final stage
FROM gcr.io/distroless/static-debian12

# Create directory for static files
WORKDIR /app/static

# Copy static files
COPY --from=builder /app/static/ ./

# Copy the binary
WORKDIR /app
COPY --from=builder /app/web-app ./

# Expose the port your app runs on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/app/web-app"]