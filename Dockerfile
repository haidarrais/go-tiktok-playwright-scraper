# # backend/Dockerfile
# FROM golang:1.23.4-bookworm AS build

# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main .


# # Final image
# FROM ubuntu:22.04

# # Install Redis client (optional, if you want to use redis-cli)
# RUN apt-get update && apt-get install -y redis-tools
# # Install Playwright and its dependencies
# RUN go install github.com/playwright-community/playwright-go/cmd/playwright@latest
# RUN playwright install --with-deps

# COPY --from=build /app/main /usr/local/bin/
# EXPOSE 3030

# CMD ["/usr/local/bin/main"]


# Stage 1: Modules caching
FROM golang:1.23.4-bookworm as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Stage 2: Build
FROM golang:1.23.4-bookworm as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /workdir
WORKDIR /workdir
# Install playwright cli with right version for later use
RUN PWGO_VER=$(grep -oE "playwright-go v\S+" /workdir/go.mod | sed 's/playwright-go //g') \
    && go install github.com/playwright-community/playwright-go/cmd/playwright@${PWGO_VER}
# Build your app
RUN GOOS=linux GOARCH=amd64 go build -o /bin/myapp

# Stage 3: Final
FROM ubuntu:noble
COPY --from=builder /go/bin/playwright /bin/myapp /
RUN apt-get update && apt-get install -y ca-certificates tzdata redis-tools \
    # Install dependencies and all browsers (or specify one)
    && /playwright install --with-deps \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/myapp /usr/local/bin/main
EXPOSE 3030

CMD ["/usr/local/bin/main"]