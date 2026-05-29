# Security Requirements

## Authentication

- Hash passwords with `bcrypt` cost 12
- Validate password length ≤ 72 bytes before hashing (bcrypt silently truncates above this)
- Access tokens: JWT, 15-minute expiry, HS256
- Refresh tokens: JWT, 7-day expiry, `httpOnly` + `Secure` + `SameSite=Strict` cookie
- Token refresh endpoint required
- Rate-limit login: 5 attempts per minute per IP

**Email verification toggle — `EMAIL_VERIFICATION_ENABLED` env var:**
- `false` (default): accounts active immediately after registration
- `true` or `1`: new accounts set `is_verified = false`; UUID token stored in `email_verifications` table with 24h expiry; verification email sent; unverified users get `403 {"error": "email_not_verified"}` on protected routes
- Always include `EMAIL_VERIFICATION_ENABLED=false` in `.env.example` with an explanatory comment

## API Security

- Validate and sanitize all user inputs — no exceptions
- Security headers on every response:
  - `X-Frame-Options: SAMEORIGIN`
  - `X-Content-Type-Options: nosniff`
  - `Referrer-Policy: strict-origin-when-cross-origin`
- CORS: allow only `CORS_ALLOWED_ORIGIN` env var — never `*` in production
- Never log passwords, tokens, or PII

**CSRF protection:**
- Enable Fiber's built-in CSRF middleware on all non-GET routes
- Double-submit cookie pattern: server sets non-httpOnly `csrf_token` cookie on login; frontend reads it and sends as `X-CSRF-Token` header on every mutation; backend validates header matches cookie

**OpenAPI / Swagger:**
- Annotate every handler with `swaggo/swag` comments
- Expose `/api/docs` (Swagger UI) only when `APP_ENV=development`
- Commit `docs/swagger.json` — regenerate after every handler change

## Docker Security

- Never run containers as root — use a non-root user in every Dockerfile
- Never put secrets in Dockerfiles — environment variables only
- Use specific image tags, never `latest`
