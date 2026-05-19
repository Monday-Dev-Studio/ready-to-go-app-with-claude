# API Contracts

Base URL: `http://localhost:8080`

All responses follow this shape:
```json
{
  "success": true,
  "data": { ... },
  "message": "optional message"
}
```

Error responses:
```json
{
  "success": false,
  "error": "error description"
}
```

---

## Auth Endpoints (Built-in)

### POST /api/auth/register
Register a new user.

**Request:**
```json
{
  "name": "Alice Smith",
  "email": "alice@example.com",
  "password": "minimum8chars"
}
```

**Response 201:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "name": "Alice Smith",
    "email": "alice@example.com",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### POST /api/auth/login
Login with email and password.

**Request:**
```json
{
  "email": "alice@example.com",
  "password": "minimum8chars"
}
```

**Response 200:**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJ...",
    "user": {
      "id": "uuid",
      "name": "Alice Smith",
      "email": "alice@example.com"
    }
  }
}
```
Refresh token is set as httpOnly cookie.

---

### POST /api/auth/refresh
Exchange refresh token (cookie) for a new access token.

**Response 200:**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJ..."
  }
}
```

---

### POST /api/auth/logout
Clear the refresh token cookie.

**Response 200:**
```json
{ "success": true, "message": "logged out" }
```

---

### GET /api/users/me
Get the current user's profile. Requires `Authorization: Bearer <token>` header.

**Response 200:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "name": "Alice Smith",
    "email": "alice@example.com",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET /health
Health check — no auth required.

**Response 200:**
```json
{ "status": "ok" }
```

---

## App Endpoints

_(Claude will add endpoints here as features are built)_
