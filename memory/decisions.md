# Architecture & Design Decisions

---

## ADR-001: Clean Architecture for Backend
**Decision:** Use clean architecture layers (domain → repository → usecase → delivery).
**Rationale:** Keeps business logic independent of frameworks and databases; makes testing easier; Claude can add features without breaking existing code.

## ADR-002: pgx over GORM
**Decision:** Use `pgx/v5` directly instead of an ORM.
**Rationale:** Explicit SQL is easier to audit for security and performance; parameterized queries prevent SQL injection by default; no magic.

## ADR-003: JWT with httpOnly Cookie for Refresh Token
**Decision:** Access token in Authorization header (short-lived, 15 min), refresh token in httpOnly cookie (7 days).
**Rationale:** Protects against XSS (refresh token is not accessible from JS); short-lived access tokens limit blast radius of token theft.

## ADR-004: React Query + Zustand
**Decision:** React Query for server state, Zustand for client state.
**Rationale:** React Query handles caching, loading, and error states for API calls. Zustand is simple and avoids Redux boilerplate for client-only state like the auth user object.

## ADR-005: Docker multi-stage builds
**Decision:** Multi-stage Docker builds for all services.
**Rationale:** Production images are small (Go binary in alpine, built React assets in nginx) with no build tools in the final image.

---

_(Claude will add new decisions here as the app grows)_
