# Create first stage docker image
FROM golang:1.22.5-alpine AS builder

WORKDIR /build

# Copy the project
COPY . .

# Get dependencies
RUN go mod tidy

# Build the app
RUN go build -C ./cmd/main -o /build/app

# Create second stage docker image
FROM scratch
WORKDIR /server

# Copy the app
COPY --from=builder /build/app ./

# Copy static files
COPY public ./

# Run the app
CMD ["/server/app"]