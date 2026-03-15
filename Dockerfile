# STAGE 1: Build the binary
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Optimization: Only download dependencies if go.mod exists
COPY go.mod ./
# If you have a go.sum file, uncomment the next line:
# COPY go.sum ./ 
RUN go mod download 

COPY . .

# Optimization: Added -ldflags="-s -w" to strip debug information 
# This makes the binary (and the final image) even smaller.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go-calculator main.go

# STAGE 2: Final Runtime
FROM alpine:latest

# Security Best Practice: Run as a non-root user
RUN adduser -D calcuser
USER calcuser

COPY --from=builder /go-calculator /usr/local/bin/calculator

# This ensures that when the container starts, it's ready for input
ENTRYPOINT ["calculator"]