# Backend Standards

## Architecture — Clean Architecture (Go-Fiber)

```
backend/
├── cmd/main.go                  # Entry point — wires all layers together
├── internal/
│   ├── domain/                  # Pure Go entities, interfaces, domain errors — no external imports
│   ├── repository/              # pgx implementations of domain interfaces
│   ├── usecase/                 # Business logic — depends on domain interfaces only
│   └── delivery/http/
│       ├── router.go            # Route registration and middleware wiring
│       └── handler/             # Thin handlers: parse → validate → call usecase → respond
└── pkg/
    ├── config/                  # Env-based config — panic on missing required vars
    ├── database/                # pgxpool setup, auto-migration runner
    ├── middleware/              # Auth (JWT), CSRF, audit log, rate limiter
    └── response/                # Standardized JSON envelope helpers
```

**Layer rules — never break these:**
- Domain imports nothing external
- Repository imports domain only
- Usecase imports repository interfaces only — never implementations
- Delivery imports usecase interfaces only
- All wiring in `cmd/main.go`

---

## Go Coding Standards

- `context.Context` as first parameter on all functions that do I/O
- Handle ALL errors — no `_` discards in production code
- Structured logging with `log/slog` — always include a `"component"` field
- `github.com/google/uuid` for UUIDs
- `github.com/go-playground/validator/v10` for input validation
- `github.com/jackc/pgx/v5` for PostgreSQL — parameterized queries only
- No raw SQL string concatenation — ever
- Handlers stay thin: parse → validate → call usecase → respond
- Return typed domain errors from usecase; map to HTTP status in handler only

**Inline comment rule:** One line, WHY only, non-obvious only. Never describe what the code does.

```go
// bcrypt silently truncates passwords longer than 72 bytes — validate length before hashing
if len(input.Password) > 72 {
    return domain.ErrPasswordTooLong
}
```

---

## SQL Standards

- Parameterized queries only: `$1, $2` placeholders
- Transactions for all multi-table writes
- `ON CONFLICT` clauses where appropriate
- `NOT NULL` by default — nullable only when truly optional

**Soft deletes — every entity table:**
- Column: `deleted_at TIMESTAMPTZ`
- Never hard-delete rows
- Every query filters `WHERE deleted_at IS NULL`

**Pagination — all list endpoints:**
- Cursor-based pagination only
- Response envelope includes `next_cursor` (string, nullable) and `has_more` (bool)

**Audit log:**
- Table `audit_logs` created in initial migrations
- Columns: `id`, `user_id`, `action`, `entity_type`, `entity_id`, `metadata JSONB`, `ip_address`, `created_at`
- Middleware logs every `POST`, `PUT`, `PATCH`, `DELETE` request automatically

**Every table gets these columns — no exceptions:**
```sql
id         UUID        PRIMARY KEY DEFAULT uuid_generate_v4()
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
deleted_at TIMESTAMPTZ
```

**Migration rules:**
- Files named `001_description.sql`, `002_description.sql` — sequential, never skip numbers
- Never modify an existing migration — always add a new one
- Every migration must be additive: `ADD COLUMN` with a default, never bare `DROP COLUMN`
- First line of every migration: `-- Migration: <what this does and why>`
- Add indexes for all foreign keys and frequently-queried columns

---

## Docker — Backend

- Development: `Dockerfile.dev` — `golang:1.24-alpine` with Air hot-reload, polling mode
- Production: multi-stage `Dockerfile` — `golang:1.24-alpine` builder → `alpine:3.21` runner
  - `CGO_ENABLED=0` for a fully static binary
  - Strip debug symbols: `-ldflags="-s -w"`
  - Run as non-root user
  - Health check via `wget /health`
