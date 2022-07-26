FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download
RUN go mod verify

# Install the package
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/backend

# Strip debugging symbols
RUN strip app

FROM --platform=linux/amd64 alpine:latest AS helper

RUN apk update

# Compress executable
WORKDIR /app
COPY --from=builder /app/app /app/toOptimize
RUN apk add --no-cache upx
RUN upx --best --lzma -o app toOptimize

# Add certs
RUN apk add ca-certificates

FROM busybox:latest AS runtime

WORKDIR /app

# Copy root certificates
COPY --from=helper /etc/ssl/certs /etc/ssl/certs

# Copy executabel files with configuration
COPY --from=helper /app/app /app/app
COPY --from=builder /app/.env /app/.env

CMD ["./app"]