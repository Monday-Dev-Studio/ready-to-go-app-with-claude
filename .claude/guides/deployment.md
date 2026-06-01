# VPS Deployment Guide

Use this guide when the user is ready to deploy to a live server. Ask for their **server IP** and **domain name** before starting. The server is always Linux — commands in Steps 2–8 are the same regardless of whether the user is on Windows or Mac.

---

## Step 1 — Connect to the Server

**Mac (Terminal):**
```bash
ssh root@YOUR_SERVER_IP
```

**Windows (PowerShell):**
```powershell
ssh root@YOUR_SERVER_IP
```

SSH works natively in PowerShell on Windows 10 and 11 — no extra tools needed.

When prompted "Are you sure you want to continue connecting?" — type `yes`.

---

## Step 2 — Server Setup (run once)

```bash
# Install Docker
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER
newgrp docker

# Install Certbot for SSL
sudo apt update && sudo apt install certbot -y

# Create backup directory
sudo mkdir -p /backups

# Open firewall
sudo ufw allow 22 && sudo ufw allow 80 && sudo ufw allow 443 && sudo ufw enable
```

---

## Step 3 — Copy the Project to the Server

**Recommended — use Git (works the same on Windows and Mac):**
```bash
# On the server
git clone https://github.com/your-username/your-repo.git /app
```

**Alternative — copy files directly:**

Mac:
```bash
scp -r /path/to/your/project root@YOUR_SERVER_IP:/app
```

Windows (PowerShell):
```powershell
scp -r C:\path\to\your\project root@YOUR_SERVER_IP:/app
```

---

## Step 4 — Get an SSL Certificate (run once per domain)

```bash
# On the server — stop anything on port 80 first
docker compose down

sudo certbot certonly --standalone -d yourdomain.com -d www.yourdomain.com
# Certificates land at: /etc/letsencrypt/live/yourdomain.com/
```

---

## Step 5 — Configure nginx

`nginx/nginx.prod.conf` must:
- Redirect all HTTP → HTTPS (301)
- Terminate TLS using certs from `/etc/letsencrypt/live/yourdomain.com/`
- Proxy `/api/` → `backend:8080`
- Serve the React SPA for all other routes: `try_files $uri /index.html`
- Set `client_max_body_size` (10m is a sensible default)
- Enable gzip for text assets (html, css, js, json, svg)

The `nginx/certs/` directory is gitignored — certs always live on the server at `/etc/letsencrypt/live/`.

---

## Step 6 — Configure Environment and Deploy

```bash
cd /app

# Fill in all production values
cp .env.example .env
nano .env

# Build and start
docker compose -f docker-compose.prod.yml up -d --build

# Watch logs to confirm everything started cleanly
docker compose -f docker-compose.prod.yml logs -f
```

Checklist before going live:
- [ ] `APP_ENV=production`
- [ ] `APP_DOMAIN` set to your real domain
- [ ] `JWT_ACCESS_SECRET` and `JWT_REFRESH_SECRET` changed to random 64-char strings (`openssl rand -hex 32`)
- [ ] `POSTGRES_PASSWORD` changed from the default
- [ ] `CORS_ALLOWED_ORIGIN` set to `https://yourdomain.com`

---

## Step 7 — Redeploy After Updates

**Mac:**
```bash
cd /app && git pull
docker compose -f docker-compose.prod.yml up -d --build --no-deps backend frontend
```

**Windows (PowerShell — run on the server via SSH):**
```powershell
# SSH into server first, then run the same Linux commands
ssh root@YOUR_SERVER_IP
cd /app && git pull
docker compose -f docker-compose.prod.yml up -d --build --no-deps backend frontend
```

---

## Step 8 — Automate SSL Renewal and Backups (run once on server)

```bash
# SSL renewal at 3am on the 1st of each month
(crontab -l 2>/dev/null; echo "0 3 1 * * certbot renew --quiet && docker compose -f /app/docker-compose.prod.yml restart nginx") | crontab -

# Daily DB backup at 2am — keep 7 days
(crontab -l 2>/dev/null; echo "0 2 * * * docker exec postgres pg_dump -U appuser appdb | gzip > /backups/db_\$(date +\%Y\%m\%d).sql.gz && find /backups -name 'db_*.sql.gz' -mtime +7 -delete") | crontab -
```

---

## Deployment Rules

- Generate `docs/DEPLOYMENT.md` on first deploy with the exact domain, server IP (redacted), and steps used
- Never hardcode the domain — always use `APP_DOMAIN` env var
- Never commit a `.env` file with real credentials — `.env.example` only
- The `nginx/certs/` directory is gitignored by design
