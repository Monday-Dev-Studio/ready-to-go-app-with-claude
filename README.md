# Ready-to-Go App Builder with Claude

Build a complete web application just by describing what you want — no coding experience needed.

This repo gives you everything you need to go from idea to running app. Claude (an AI assistant) writes all the code. Docker runs it. You just need to describe your vision.

---

## What You Will Get

- A full web application with a backend API, frontend website, and database
- Secure user login and registration (out of the box)
- Runs entirely in Docker — no need to install Go, Node.js, or PostgreSQL
- A technical document that explains your app to any developer
- A memory system so Claude never forgets what you built

---

## What You Need (Prerequisites)

You only need two things installed on your computer:

### 1. Claude Code
Claude Code is the AI assistant that writes all the code for you.

**Mac:**
```
brew install claude
```
Or download from: https://claude.ai/download

**Windows:**
Download the installer from: https://claude.ai/download

> After installing, run `claude --version` in your terminal to confirm it works.

### 2. Docker Desktop
Docker runs your application without you needing to install Go, Node.js, or PostgreSQL.

**Download Docker Desktop:** https://www.docker.com/products/docker-desktop/

- Mac: Download the `.dmg` file and drag to Applications
- Windows: Download the `.exe` installer and follow the wizard

> After installing, open Docker Desktop and wait for the whale icon to stop animating. Then run `docker --version` to confirm.

---

## Step 1 — Clone This Repository

Open your terminal (Mac: search "Terminal", Windows: search "Command Prompt" or "PowerShell") and run:

```bash
git clone https://github.com/Monday-Dev-Studio/ready-to-go-app-with-claude.git
cd ready-to-go-app-with-claude
```

---

## Step 2 — Set Up Your Environment File

Copy the example environment file:

**Mac / Linux:**
```bash
cp .env.example .env
```

**Windows (Command Prompt):**
```cmd
copy .env.example .env
```

**Windows (PowerShell):**
```powershell
Copy-Item .env.example .env
```

You can leave the values as-is for local development. The defaults work out of the box.

---

## Step 3 — Start Claude Code

Inside the project folder, start Claude Code:

```bash
claude
```

You will see a prompt where you can type messages to Claude.

---

## Step 4 — Tell Claude What You Want to Build

Just describe your app idea in plain English! Here are some example prompts to get you started:

---

### Example 1 — Simple task manager

```
I want to build a task management app. Users can register and log in. 
After logging in, they can create tasks with a title and description, 
mark tasks as complete, and delete tasks. Show all tasks in a list 
sorted by creation date (newest first).
```

---

### Example 2 — Blog platform

```
Build me a blog platform. There are two types of users: admins and readers. 
Admins can write, edit, and delete blog posts with a title, content, and 
cover image URL. Readers can read posts and leave comments. 
Show posts in reverse chronological order on the home page.
```

---

### Example 3 — Inventory tracker

```
I need an inventory management system for a small shop. Users can add products 
with a name, SKU, quantity, and price. They can update the quantity when stock 
changes. Show a dashboard with total products, total stock value, and 
products that are low on stock (less than 10 units).
```

---

Claude will ask you a few clarifying questions, then start building your app. This may take several minutes. Claude will tell you when it is done.

---

## Step 5 — Run Your App

Once Claude finishes building, start your application with Docker:

```bash
docker compose up --build
```

This command:
- Builds your Go backend
- Builds your React frontend
- Starts PostgreSQL
- Runs database migrations automatically
- Starts everything together

The first time you run this, it may take 3–5 minutes to download and build everything. After that, it will be much faster.

When you see output like this, your app is ready:

```
backend   | Server started on port 8080
frontend  | Local:   http://localhost:3000
```

Open your browser and go to: **http://localhost:3000**

---

## Windows Users — Important Setup

Before running Docker on Windows, make sure:

1. **Docker Desktop is using WSL2** (not Hyper-V)
   - Open Docker Desktop → Settings → General → check "Use the WSL 2 based engine"

2. **Run commands in PowerShell or Windows Terminal** (not Command Prompt)

3. **Clone using Git for Windows** with LF line endings:
   ```powershell
   git config --global core.autocrlf false
   git clone https://github.com/Monday-Dev-Studio/ready-to-go-app-with-claude.git
   ```
   > If you already cloned, run `git rm --cached -r .` then `git reset --hard` inside the folder to fix line endings.

4. **Copy the env file:**
   ```powershell
   Copy-Item .env.example .env
   ```

5. Then run normally:
   ```powershell
   docker compose up --build
   ```

---

## How to Stop the App

Press `Ctrl + C` in the terminal where Docker is running.

To stop and remove containers (clean state):
```bash
docker compose down
```

To also delete the database data (full reset):
```bash
docker compose down -v
```

---

## Step 6 — Adding New Features

Want to add something new? Go back to Claude and describe what you want:

```bash
claude
```

Then type your request. Examples:

```
Add the ability for users to attach files to tasks. 
Support PDF and images up to 5MB.
```

```
Add a search feature to find blog posts by title or content.
```

```
Add email notifications when a task is assigned to a user.
```

Claude will look at what was already built and add the new feature without breaking existing functionality.

After Claude is done, restart Docker to apply the changes:

```bash
docker compose up --build
```

---

## Folder Structure (for the curious)

```
ready-to-go-app-with-claude/
├── CLAUDE.md          ← Instructions for Claude (don't delete this!)
├── README.md          ← This file
├── .env               ← Your local configuration (never commit this)
├── .env.example       ← Template for configuration
├── docker-compose.yml ← Runs everything locally
├── memory/            ← Claude's memory about your project
├── docs/
│   └── TECH_DOC.md   ← Technical documentation (auto-updated by Claude)
├── backend/           ← Go-Fiber API server
└── frontend/          ← React web application
```

---

## Troubleshooting

### "docker: command not found"
Docker Desktop is not installed or not running. Download it from https://www.docker.com/products/docker-desktop/ and make sure the whale icon is visible in your taskbar/menu bar.

### "claude: command not found"
Claude Code is not installed. Follow Step 1 above.

### Port already in use
Something else is using port 3000 or 8080. Stop the other application, or edit `.env` to change the ports.

### The app won't start / shows errors
Run `docker compose logs` to see detailed error messages. Copy the error and ask Claude:
```
The app won't start. Here is the error: [paste the error here]. What is wrong?
```

### "permission denied" on Mac/Linux
Run: `chmod +x scripts/*.sh`

---

## For Developers

If you are a developer taking over this project, read `docs/TECH_DOC.md` for the full technical documentation including:
- Architecture overview
- API reference
- Database schema
- Environment variables
- Deployment guide

---

## Production Deployment

When you are ready to deploy to a real server:

```bash
docker compose -f docker-compose.prod.yml up --build -d
```

Ask Claude to help you set up deployment to any cloud provider:
```
Help me deploy this app to [DigitalOcean / AWS / Google Cloud / Railway / Render].
```

---

## Need Help?

Ask Claude anything about the app:
- "Explain how the login system works"
- "What database tables exist?"
- "How do I back up the database?"
- "Add a new feature: [describe it]"
- "Fix this bug: [describe the problem]"

Claude knows everything about your project and is always ready to help.
