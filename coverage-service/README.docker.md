# Coverage Dashboard — Docker Guide

---

## Images on Docker Hub

Repository: `debopriyalahiri/coverage-dashboard`

| Tag | Description | Updated |
|---|---|---|
| `backend-latest` | Go API server — always latest | Auto on every push to main |
| `frontend-latest` | Vue 3 dashboard via nginx — always latest | Auto on every push to main |

Pull the images:
```bash
docker pull debopriyalahiri/coverage-dashboard:backend-latest
docker pull debopriyalahiri/coverage-dashboard:frontend-latest
```

---

## Option 1 — Run from Docker Hub (for others)

No source code needed. Just three files required:

- `docker-compose.hub.yml`
- `backend/config/services.docker.json`
- `.env`

**Step 1 — Create `.env` next to `docker-compose.hub.yml`:**
```
POSTGRES_USER=coverage_user
POSTGRES_PASSWORD=coverage_pass
POSTGRES_DB=coverage_db
```

**Step 2 — Edit `services.docker.json` with your service paths:**
```json
[
  {
    "name": "my-service",
    "display_name": "My Service",
    "backend_path": "/services/my-service/backend",
    "frontend_path": "/services/my-service/frontend",
    "backend_type": "go",
    "frontend_type": "npm"
  }
]
```

**Step 3 — Edit the volume path in `docker-compose.hub.yml`:**
```yaml
volumes:
  - /your/projects/folder:/services   ← change left side only
```

Windows example: `C:/Users/YourName/poc2_ami:/services`
Mac/Linux example: `/home/yourname/poc2_ami:/services`

**Step 4 — Start:**
```bash
docker compose -f docker-compose.hub.yml up
```

Open `http://localhost:5173`

---

## Option 2 — Run from Local Build (for development)

Full source code needed. Builds images locally from source.

**Step 1 — Create `.env`:**
```
POSTGRES_USER=coverage_user
POSTGRES_PASSWORD=coverage_pass
POSTGRES_DB=coverage_db
```

**Step 2 — Edit volume path in `docker-compose.yml`:**
```yaml
volumes:
  - E:/poc2_ami:/services   ← change to your path
```

**Step 3 — First run or after code changes:**
```bash
docker compose up --build
```

**Step 4 — Subsequent runs (no code changes):**
```bash
docker compose start
```

Open `http://localhost:5173`

---

## Option 3 — Run without Docker (local Go + Node)

**Terminal 1 — Backend:**
```bash
cd coverage-service/backend
go run main.go
```

**Terminal 2 — Frontend:**
```bash
cd coverage-service/frontend
npm install
npm run dev
```

Open `http://localhost:5173`

Uses `services.json` with relative paths (`../../hello-world/backend`).

---

## Docker Commands Reference

| Command | What it does |
|---|---|
| `docker compose up --build` | First time or after code changes |
| `docker compose up` | Start everything (no rebuild) |
| `docker compose start` | Resume stopped containers (fastest) |
| `docker compose stop` | Stop but keep data |
| `docker compose down` | Stop and remove containers (data safe) |
| `docker compose down -v` | Stop and delete everything including database |
| `docker compose restart backend` | Restart only backend (e.g. after services.json change) |
| `docker compose -f docker-compose.hub.yml up` | Run using Hub images |
| `docker compose ps` | Check what is running |
| `docker compose logs -f backend` | See live backend logs |

---

## services.json — Path format by mode

| Mode | Path format |
|---|---|
| `go run` (local) | `../../hello-world/backend` |
| Docker (local build or Hub) | `/services/hello-world/backend` |

Two files kept separately:
- `backend/config/services.json` — local paths, used by `go run`
- `backend/config/services.docker.json` — `/services/` paths, used by Docker

---

## CI/CD — GitHub Actions

Every push to `main` that changes anything inside `coverage-service/` automatically:

1. Builds new backend and frontend images
2. Pushes `backend-latest` and `frontend-latest` to Docker Hub

No manual push needed after the initial setup.

**Required GitHub Secrets:**

| Secret | Value |
|---|---|
| `DOCKER_USERNAME` | `debopriyalahiri` |
| `DOCKER_TOKEN` | Docker Hub personal access token |

Workflow file: `.github/workflows/docker-publish.yml`

---

## Troubleshooting

**Path not found when running tests**
Check that the volume path in compose file matches your actual folder.
Check that `services.docker.json` paths start with `/services/`.

**Port already in use**
Change the left port number:
```yaml
- "5174:80"   
```

**Database connection error**
Wait a few seconds — backend starts before PostgreSQL is fully ready.
Run `docker compose up` again.
