# Technical Documentation

> This document is automatically maintained by Claude. It is updated after every feature addition or change.

**Stack:** Go 1.24 · Fiber v2 · React 18 · TypeScript 5 · PostgreSQL 16 · Docker
**Last Updated:** _(Claude will fill this in)_

---

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Project Structure](#project-structure)
3. [Environment Variables](#environment-variables)
4. [Database Schema](#database-schema)
5. [API Reference](#api-reference)
6. [Running Locally](#running-locally)
7. [Running in Production](#running-in-production)
8. [Testing](#testing)
9. [Changelog](#changelog)

---

## Architecture Overview

```
┌─────────────────────────────────────────────────┐
│                   Browser                        │
│              React (TypeScript)                  │
│   React Query · Zustand · react-hook-form        │
└───────────────────┬─────────────────────────────┘
                    │ HTTP (REST)
┌───────────────────▼─────────────────────────────┐
│              Go-Fiber API Server                 │
│                                                  │
│  Delivery (Handlers)                             │
│       ↓                                          │
│  Use Cases (Business Logic)                      │
│       ↓                                          │
│  Repositories (Data Access)                      │
│       ↓                                          │
│  PostgreSQL                                      │
└─────────────────────────────────────────────────┘
```

The backend follows Clean Architecture:
- **Domain** — pure Go structs, no framework dependencies
- **Repository** — implements database operations using pgx
- **Use Case** — orchestrates business logic using repository interfaces
- **Delivery** — Fiber HTTP handlers that call use cases

---

## Project Structure

```
.
├── CLAUDE.md              ← AI instructions (do not modify)
├── README.md              ← User guide
├── .env                   ← Local config (never commit)
├── .env.example           ← Config template
├── docker-compose.yml     ← Local development
├── docker-compose.prod.yml← Production deployment
├── memory/                ← Claude's project memory
└── docs/
    └── TECH_DOC.md        ← This file

backend/
├── cmd/main.go            ← Entry point, dependency injection
├── internal/
│   ├── domain/            ← Entities and domain errors
│   ├── repository/        ← DB interfaces + pgx implementations
│   ├── usecase/           ← Business logic
│   └── delivery/http/     ← Fiber routes and handlers
└── pkg/
    ├── config/            ← Environment config loader
    ├── database/          ← PostgreSQL connection pool
    ├── middleware/         ← JWT auth, CORS, rate limiter
    └── response/          ← Standard JSON response helpers

frontend/
└── src/
    ├── pages/             ← Page-level components (one per route)
    ├── components/        ← Reusable UI components
    ├── hooks/             ← Custom React hooks
    ├── services/          ← Axios-based API functions
    ├── store/             ← Zustand state stores
    └── types/             ← TypeScript interfaces
```

---

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | `development` or `production` | `development` |
| `APP_PORT` | Backend port | `8080` |
| `FRONTEND_PORT` | Frontend port | `3000` |
| `POSTGRES_HOST` | Database host | `postgres` |
| `POSTGRES_PORT` | Database port | `5432` |
| `POSTGRES_USER` | Database user | `appuser` |
| `POSTGRES_PASSWORD` | Database password | _(must set)_ |
| `POSTGRES_DB` | Database name | `appdb` |
| `JWT_ACCESS_SECRET` | Secret for access tokens | _(must set)_ |
| `JWT_REFRESH_SECRET` | Secret for refresh tokens | _(must set)_ |
| `JWT_ACCESS_EXPIRY_MINUTES` | Access token lifetime | `15` |
| `JWT_REFRESH_EXPIRY_DAYS` | Refresh token lifetime | `7` |
| `CORS_ALLOWED_ORIGIN` | Frontend URL for CORS | `http://localhost:3000` |

---

## Database Schema

See `memory/database_schema.md` for the full schema.

---

## API Reference

See `memory/api_contracts.md` for the full API reference.

---

## Running Locally

**Prerequisites:** Docker Desktop

```bash
# 1. Copy env file
cp .env.example .env

# 2. Start all services (builds images on first run)
docker compose up --build

# 3. Open the app
open http://localhost:3000
```

**Useful commands:**

```bash
# View logs
docker compose logs -f

# View only backend logs
docker compose logs -f backend

# Reset database (deletes all data)
docker compose down -v && docker compose up --build

# Rebuild a single service
docker compose up --build backend
```

---

## Running in Production

```bash
# Copy and configure env for production
cp .env.example .env
# Edit .env: set strong passwords, JWT secrets, CORS origin

# Build and start in background
docker compose -f docker-compose.prod.yml up --build -d

# View logs
docker compose -f docker-compose.prod.yml logs -f

# Update to latest code
git pull
docker compose -f docker-compose.prod.yml up --build -d
```

---

## Testing

```bash
# Run backend unit tests
docker compose exec backend go test ./...

# Run with coverage
docker compose exec backend go test -cover ./...
```

---

## Changelog

_(Claude will add entries here as features are built)_

### Initial Setup
- Go-Fiber backend with clean architecture scaffolding
- React (TypeScript) frontend with routing and auth
- PostgreSQL with migration runner
- JWT authentication (register, login, refresh, logout)
- Docker Compose for local development and production
