# POC — Test Coverage Dashboard

A monorepo containing two sample services and a standalone coverage dashboard that automatically runs their tests and displays results in a live web UI.

---

## Project Structure

```
poc2_ami/
├── coverage-service/        ← standalone coverage dashboard (this is the main app)
│   ├── backend/             ← Go API server (port 8081)
│   ├── frontend/            ← Vue 3 dashboard (port 5173)
│   └── docker-compose.yml  ← starts everything with one command
│
├── hello-world/             ← sample service 1
│   ├── backend/             ← Go REST API (port 8082)
│   └── frontend/            ← Vue 3 app (port 5173)
│
└── inventory-tracker/       ← sample service 2
    ├── backend/             ← Go REST API
    └── frontend/            ← Vue app
```

---

## Quick Start

### Option 1 — Docker (recommended)

> Requires Docker Desktop installed. No Go or Node needed.

**1. Edit the volume path in `coverage-service/docker-compose.yml`:**

```yaml
volumes:
  - E:/poc2_ami:/services   # ← change to your actual path
```

**2. Create `.env` in `coverage-service/`:**

```
POSTGRES_USER=coverage_user
POSTGRES_PASSWORD=coverage_pass
POSTGRES_DB=coverage_db
```

**3. Start everything:**

```bash
cd coverage-service
docker compose up --build
```

**4. Open the dashboard:**

```
http://localhost:5173
```

---

### Option 2 — Local (without Docker)

> Requires Go 1.21+, Node 18+, PostgreSQL running locally.

**Terminal 1 — Coverage backend:**

```bash
cd coverage-service/backend
go run main.go
```

**Terminal 2 — Coverage frontend:**

```bash
cd coverage-service/frontend
npm install
npm run dev
```

Open `http://localhost:5173`

---

## How to Use the Dashboard

1. Open `http://localhost:5173`
2. Select a service from the dropdown
3. Click **Run Tests**
4. Watch test output stream live in the terminal
5. Results save automatically to the database and appear in the table

---

## Adding a New Service

Edit `coverage-service/backend/config/services.json`:

```json
{
  "name": "my-service",
  "display_name": "My Service",
  "backend_path": "/services/my-service/backend",
  "frontend_path": "/services/my-service/frontend",
  "backend_type": "go",
  "frontend_type": "npm"
}
```

For local mode use relative paths (`../../my-service/backend`).
Then restart the backend — the service appears in the dropdown automatically.

---

## Docker Commands

| Command | What it does |
|---|---|
| `docker compose up --build` | First time or after code changes |
| `docker compose up` | Start everything (no rebuild) |
| `docker compose start` | Resume stopped containers (fastest) |
| `docker compose stop` | Stop but keep data |
| `docker compose down` | Stop and remove containers (data safe) |
| `docker compose down -v` | Stop and delete everything including database |
| `docker compose restart backend` | Restart only backend (e.g. after services.json change) |
| `docker compose ps` | Check what is running |
| `docker compose logs -f backend` | See live backend logs |

---

## Stack

| Layer | Tech |
|---|---|
| Coverage Backend | Go · Gin · GORM |
| Coverage Frontend | Vue 3 · TypeScript · Vite |
| Database | PostgreSQL |
| Containerisation | Docker · Docker Compose |
| Service backends | Go |
| Service frontends | Vue 3 · TypeScript · Vite |
| Testing | go test · Vitest |

---

## Services

### Hello World
A minimal Go + Vue service used to demonstrate the coverage dashboard.
- Backend runs on port `8082`
- Frontend runs on port `5173`

### Inventory Tracker
A CRUD inventory management service with authentication.
- Backend runs on port defined in its `.env`
- Frontend runs on port `5173`

---

## Notes

- Never commit `.env` files — they are in `.gitignore`
- `services.json` has two versions:
  - `services.json` — local paths for `go run`
  - `services_docker.json` — `/services/` paths for Docker (mounted via volume)
- The coverage image does not need rebuilding when `services.json` changes — just restart the backend container
