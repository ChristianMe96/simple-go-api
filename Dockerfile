FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY main.go .

# Build the application
# CGO_ENABLED=0 for static binary, GOOS=linux for Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple-go-api .

FROM alpine:latest

# Install ca-certificates (useful for HTTPS requests if needed)
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN adduser -D -s /bin/sh appuser

WORKDIR /app

COPY --from=builder /app/simple-go-api .

# Change ownership to non-root user
RUN chown appuser:appuser /app/simple-go-api

USER appuser

CMD ["./simple-go-api"]