# Hello World — Vue 3 + Go

A simple full-stack Hello World app with:
- **Frontend**: Vue 3 + TypeScript + Vite (port 5173)
- **Backend**: Go HTTP server (port 8082)
- **Tests**: Vitest (FE) + Go testing (BE)

---

## Quick Start

### Backend (Go)

```bash
cd backend
go run main.go
```

Runs on http://localhost:8082

**Run tests:**
```bash
cd backend
go test ./...
```

---

### Frontend (Vue)

```bash
cd frontend
npm install
npm run dev
```

Runs on http://localhost:5173  
Proxies `/api/*` → `http://localhost:8082`

**Run tests:**
```bash
cd frontend
npm test
```

---

## API Endpoints

| Method | Path          | Description        |
|--------|---------------|--------------------|
| GET    | /api/hello    | Returns Hello World JSON |
| GET    | /api/health   | Health check       |

### Example response — `/api/hello`

```json
{
  "message": "Hello, World!",
  "status": "ok"
}
```

---

## Project Structure

```
hello-world/
├── backend/
│   ├── main.go          # HTTP server + handlers
│   ├── main_test.go     # Go tests
│   └── go.mod
└── frontend/
    ├── src/
    │   ├── components/
    │   │   └── HelloWorld.vue   # Main component
    │   ├── __tests__/
    │   │   └── HelloWorld.test.ts  # Vitest tests
    │   ├── App.vue
    │   └── main.ts
    ├── index.html
    ├── vite.config.ts
    └── package.json
```
