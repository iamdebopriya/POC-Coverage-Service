# Coverage Service — Docker Setup

## What you need installed
- Docker Desktop (Windows / Mac) or Docker Engine (Linux)
- That's it. No Go, Node, or PostgreSQL needed on your machine.

---

## Step 1 — Edit docker-compose.yml (ONE line)

Open `docker-compose.yml` and find the volumes section under `backend`.
Change the left side to point to your root folder:

```yaml
volumes:
  - C:/Users/YourName/poc2_ami:/services   # ← change this path
```

**Windows:**  `C:/Users/HP/poc2_ami:/services`
**Mac/Linux:** `/home/yourname/poc2_ami:/services`

The right side (`:/services`) must stay as-is.

---

## Step 2 — Edit services.json (if your services have different names)

Open `backend/config/services.json`.
Paths must start with `/services/` — that is where Docker mounts your folder.

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

---

## Step 3 — Start everything

```bash
docker compose up --build
```

First run takes 2–3 minutes (downloads images, builds).
After that, subsequent starts are fast.

---

## Step 4 — Open the dashboard

```
http://localhost:5173
```

Pick a service → click Run Tests → watch output stream live.

---

## Useful commands

```bash
# Start (foreground, see logs)
docker compose up --build

# Start in background
docker compose up --build -d

# Stop
docker compose down

# Stop and delete database too
docker compose down -v

# See logs
docker compose logs -f backend
docker compose logs -f frontend

# Rebuild after code changes
docker compose up --build
```

---

## Ports used

| Service    | Port |
|------------|------|
| Frontend   | 5173 |
| Backend    | 8081 |
| PostgreSQL | 5432 |

---

## Folder structure expected

```
poc2_ami/                   ← your root (mounted as /services inside Docker)
├── coverage-service/       ← this app
├── hello-world/
│   ├── backend/
│   └── frontend/
└── inventory-tracker/
    ├── backend/
    └── frontend/
```

---

## Docker Hub

Images available at:
https://hub.docker.com/r/debopriyalahiri/coverage-dashboard

| Tag | Description |
|---|---|
| `backend-1.0.0` | Go API server |
| `frontend-1.0.0` | Vue 3 dashboard |

---

## Troubleshooting

**"Cannot connect to database"**
Wait a few seconds — the backend starts before PostgreSQL is fully ready.
Run `docker compose up` again.

**"Path not found" when running tests**
Check that the volume path in `docker-compose.yml` matches your actual folder location.
Check that `services.json` paths start with `/services/`.

**Port already in use**
Change the left side of the port mapping:
`"5174:80"` instead of `"5173:80"` for the frontend.

docker compose up

