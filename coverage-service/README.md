# Coverage Service — Full UI Flow

## How it works

1. You register your services in `backend/config/services.json`
2. Start the coverage backend + frontend
3. Open the dashboard → pick a service from the dropdown → click Run Tests
4. Tests run live — you see output streaming in the terminal
5. When done, results auto-save to DB and appear in the table

---

## Folder layout expected

Your projects must sit next to coverage-service:

```
projects/
├── coverage-service/     ← this repo
├── hello-world/
│   ├── backend/          ← go test runs here
│   └── frontend/         ← npm test runs here
└── inventory-tracker/
    ├── backend/
    └── frontend/
```

If your layout is different, update the paths in:
`backend/config/services.json`

---

## Step 1 — Add your services to services.json

File: `backend/config/services.json`

```json
[
  {
    "name": "hello-world",
    "display_name": "Hello World",
    "backend_path": "../hello-world/backend",
    "frontend_path": "../hello-world/frontend",
    "backend_type": "go",
    "frontend_type": "npm"
  },
  {
    "name": "inventory-tracker",
    "display_name": "Inventory Tracker",
    "backend_path": "../inventory-tracker/backend",
    "frontend_path": "../inventory-tracker/frontend",
    "backend_type": "go",
    "frontend_type": "npm"
  }
]
```

Paths are relative to the `backend/` folder.

---

## Step 2 — Start the backend

```bash
cd coverage-service/backend
go run main.go
```

Runs on http://localhost:8081

---

## Step 3 — Start the frontend

```bash
cd coverage-service/frontend
npm install
npm run dev
```

Open http://localhost:5173

---

## Step 4 — Run tests from the dashboard

1. Open http://localhost:5173
2. In the **Run Tests** panel, pick a service from the dropdown
3. Click **Run Tests**
4. Watch the terminal — output streams live
5. When finished, the table updates automatically

---

## API reference

| Method | Path                        | What it does                        |
|--------|-----------------------------|-------------------------------------|
| GET    | /api/registered-services    | List services from services.json    |
| GET    | /api/run/:service           | Run tests (SSE stream)              |
| POST   | /api/coverage               | Save a result                       |
| GET    | /api/coverage               | List results (?service=&from=&to=)  |
| GET    | /api/coverage/services      | Services that have results in DB    |
| GET    | /api/coverage/download      | Download CSV                        |
