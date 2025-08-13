# ---- Build stage ----
FROM golang:1.24-alpine AS builder
WORKDIR /src

# Optional (some deps need git)
RUN apk add --no-cache git

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources
COPY . .

# Build static-ish binary (no DWARF, smaller)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/app ./cmd

# ---- Runtime stage ----
FROM alpine:3.20 AS runner
# For HTTPS calls inside your app
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /out/app /app/app

EXPOSE 5000
CMD ["/app/app"]
