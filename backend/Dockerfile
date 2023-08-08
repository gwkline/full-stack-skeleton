# Build stage
FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest

# If you need bash or openssh at runtime, install them here.
# Otherwise, omit this.
RUN apk --no-cache add bash

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .
COPY --from=builder /app/wait.sh .

# Make sure the script is executable
RUN chmod +x ./wait.sh

EXPOSE 8888

CMD ./wait.sh database:5432 -- ./main