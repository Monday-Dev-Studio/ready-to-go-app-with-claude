# Docker Setup & Usage Guide

Use this guide when a user needs to install Docker or learn how to use it. In Non-Tech Mode, walk through it step by step. In Tech Mode, point to the relevant section.

---

## What is Docker? (Non-Tech explanation)

> Docker is like a self-contained box that holds your entire app — the database, the server, and the website — all pre-configured. You press one command and everything starts. You don't need to install Go, Node.js, or PostgreSQL on your computer. The box handles all of that.

---

## Step 1 — Install Docker Desktop

### Windows

1. Go to [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/) and download **Docker Desktop for Windows**
2. Run the installer — it will ask to install **WSL 2** (Windows Subsystem for Linux). Click **Yes** — this is required
3. Restart your computer when prompted
4. Open **Docker Desktop** from the Start menu
5. Wait for the whale icon in the taskbar to stop animating — that means Docker is ready

**If you see "WSL 2 installation is incomplete":**
Open PowerShell as Administrator and run:
```powershell
wsl --install
```
Restart your computer and try again.

### Mac

1. Go to [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/) and download Docker Desktop for Mac
   - M1/M2/M3 Mac → download **Apple Silicon** version
   - Older Mac → download **Intel Chip** version
2. Open the downloaded `.dmg` file and drag Docker to your Applications folder
3. Open Docker from Applications
4. Wait for the whale icon in the menu bar to stop animating — Docker is ready

---

## Step 2 — Verify Installation

Open your terminal (PowerShell on Windows, Terminal on Mac) and run:

```
docker --version
docker compose version
```

Both should print a version number. If they do, Docker is ready to use.

---

## Step 3 — Essential Commands

| Command | What it does |
|---------|-------------|
| `docker compose up` | Start everything (database, backend, frontend) — logs stream in the terminal |
| `docker compose up -d` | Start everything in the background (no log output) |
| `docker compose down` | Stop everything cleanly |
| `docker compose down -v` | Stop everything AND wipe all data — caution, deletes the database |
| `docker compose logs -f` | Watch live logs from all services |
| `docker compose logs -f backend` | Watch live logs from the backend only |
| `docker compose ps` | See which services are running |
| `docker compose build` | Rebuild images (run after changing Dockerfiles) |
| `docker ps` | See all running containers on your machine |
| `docker system prune` | Free up disk space — removes unused images and containers |

All of these commands work the same on Windows (PowerShell) and Mac (Terminal).

---

## Step 4 — Starting the Development App

```
docker compose up
```

Then open your browser:
- **App:** `http://localhost:3000`
- **API:** `http://localhost:8080`
- **API Docs:** `http://localhost:8080/api/docs` (development only)

Hot-reload is active — save a file and the app updates automatically.

---

## Common Problems and Fixes

**"Port already in use" / "address already in use"**
Another app is using the same port. Run `docker ps` to see running containers, then `docker compose down` to stop them.

**"Cannot connect to the Docker daemon"**
Docker Desktop isn't running. Open it from Start menu (Windows) or Applications (Mac) and wait for it to be fully ready.

**Windows: "WSL 2 installation is incomplete"**
Open PowerShell as Administrator and run `wsl --install`. Restart and try again.

**Windows: files not hot-reloading**
Polling mode is already configured in this repo. If it still doesn't reload, run `docker compose down` then `docker compose up` again.

**"No space left on device"**
Run `docker system prune` to free up disk space, then try again.

**Database won't start / keeps restarting**
Run `docker compose down -v` to clear old database data, then `docker compose up` again. This resets all local data.

**Changes not showing up after code edit**
Make sure you saved the file. If the issue persists, run `docker compose restart backend` (or `frontend`) to force a restart.
