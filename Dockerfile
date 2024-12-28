# backend/Dockerfile
FROM golang:1.23.4-bookworm AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Final image
FROM ubuntu:22.04

# Install Redis client (optional, if you want to use redis-cli)
RUN apt-get update && apt-get install -y redis-tools

COPY --from=build /app/main /usr/local/bin/
EXPOSE 3030

CMD ["/usr/local/bin/main"]