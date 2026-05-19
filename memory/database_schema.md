# Database Schema

Database: PostgreSQL 16
Extensions: `uuid-ossp`

---

## users

| Column        | Type                     | Constraints              |
|---------------|--------------------------|--------------------------|
| id            | UUID                     | PRIMARY KEY, default uuid_generate_v4() |
| email         | VARCHAR(255)             | UNIQUE NOT NULL          |
| password_hash | VARCHAR(255)             | NOT NULL                 |
| name          | VARCHAR(255)             | NOT NULL                 |
| created_at    | TIMESTAMP WITH TIME ZONE | DEFAULT NOW()            |
| updated_at    | TIMESTAMP WITH TIME ZONE | DEFAULT NOW()            |

**Indexes:** `idx_users_email` on `(email)`

---

## App Tables

_(Claude will add tables here as features are built)_
