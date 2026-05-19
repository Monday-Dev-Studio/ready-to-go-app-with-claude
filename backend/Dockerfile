# ─── Build stage ───────────────────────────────────────────────────────────────
FROM golang:1.24-alpine AS builder

WORKDIR /build

# Cache dependencies before copying source
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/main.go

# ─── Production stage ──────────────────────────────────────────────────────────
FROM alpine:3.21 AS production

RUN apk add --no-cache ca-certificates wget

# Non-root user for security
RUN addgroup -S app && adduser -S app -G app
USER app

WORKDIR /app

COPY --from=builder /build/server .
COPY --chown=app:app migrations/ ./migrations/

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=5s --start-period=30s --retries=5 \
    CMD wget -qO- http://localhost:8080/health || exit 1

CMD ["./server"]
