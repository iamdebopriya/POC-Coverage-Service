# Coverage Service — Full UI Flow

## Overview

This repository provides a dashboard for running backend and frontend tests across multiple projects, streaming test output live, and saving coverage + result summaries into a database.

The dashboard is built as a separate service that discovers your registered applications from `backend/config/services.json` and uses HTTP + Server-Sent Events to deliver live terminal output while tests run.

## Core files and folders

### Backend

- `backend/main.go`
  - Entry point for the Go server.
  - Loads environment variables from `.env` and starts the Gin HTTP server.
  - Connects to the Postgres database and auto-migrates the `Coverage` model.

- `backend/router/router.go`
  - Configures API routes and CORS headers.
  - Exposes endpoints for service registration, coverage CRUD, report downloads, and running tests.

- `backend/config/config.go`
  - Reads `backend/config/services.json` into `ServiceConfig` objects.
  - Uses relative paths from `backend/` to locate each target project's `backend` and `frontend` directories.

- `backend/config/services.json`
  - Contains the list of services the dashboard can run.
  - Each service includes `name`, `display_name`, `backend_path`, `frontend_path`, `backend_type`, and `frontend_type`.

- `backend/handlers/run_script.go`
  - Core test runner logic.
  - Implements `RunTestsSSE`, which instantiates an SSE stream and executes configured test commands.
  - Parses command output to derive pass/fail counts, coverage percentages, and flaky test counts.
  - Saves results to the database after the run completes.

- `backend/handlers/coverage.go`
  - CRUD endpoints for coverage history.
  - Filter by service, date range, and timezone offset.
  - Exports results to CSV.

- `backend/database/db.go`
  - Connects to Postgres using environment variables.
  - Uses GORM as the ORM layer.

- `backend/models/coverage.go`
  - Defines the `Coverage` schema stored in the DB.
  - Tracks service name, backend/frontend coverage, tests passed/failed/flaky, average execution time, and timestamp.

### Frontend

- `frontend/src/api.ts`
  - Contains API helper functions for fetching registered services, coverage history, and running tests.
  - Implements `runServiceTests()` using `EventSource` to receive SSE updates from `/api/run/:service`.

- `frontend/src/views/Dashboard.vue`
  - Main dashboard page.
  - Loads services and coverage history on mount.
  - Displays stats, latest run information, filters, and history table.
  - Re-fetches coverage data after a run completes.

- `frontend/src/components/RunPanel.vue`
  - Service selection and live terminal panel.
  - Starts/stops test runs and renders SSE lines.
  - Emits `done` when the run completes so parent view can refresh history.

- `frontend/src/components/FilterBar.vue`
  - Filter UI for service, date range, and CSV export.
  - Sends filter updates back to `Dashboard.vue`.

- `frontend/src/components/CoverageTable.vue`
  - Displays saved coverage rows with pass rate, coverage, and timestamp.

## Detailed flow: test running and result saving

### 1. User action

- The user opens the dashboard at `http://localhost:5173`.
- They select a service from the `Run Tests` dropdown and click **Run Tests**.

### 2. Frontend sends an SSE request

- `RunPanel.vue` calls `runServiceTests(serviceName, onLine, onDone, onError)`.
- `runServiceTests()` opens an `EventSource` to `/api/run/:service`.
- Each incoming SSE message is appended to the terminal pane in real time.

### 3. Backend receives `/api/run/:service`

- `backend/handlers/run_script.go` looks up the service in `registeredServices` loaded from `backend/config/services.json`.
- If the service is found, it starts streaming output with SSE headers.
- The call is long-lived until all configured tests finish.

### 4. Running backend tests

- If the service config defines `backend_type: "go"`, the backend runs:
  - `go test ./... -v -coverpkg=./... -cover`
- It resolves `backend_path` relative to the `backend/` directory.
- Output is streamed line-by-line as SSE events.

### 5. Running frontend tests

- If the service config defines `frontend_type: "npm"`, the backend runs:
  - `npm install --silent`
  - `npm test -- --coverage`
- This ensures dependencies are installed before the test run.
- Output is streamed live to the dashboard terminal.

### 6. Parsing output

- `runCommand()` captures both `stdout` and `stderr`.
- It strips ANSI color codes and scans each line.
- It extracts:
  - coverage percentages from Go and Vitest output
  - pass/fail counts for Go tests
  - Vitest pass counts
  - flaky tests by detecting tests that both pass and fail in the same run
- The parsed metrics are accumulated into a `models.Coverage` result.

### 7. Summary and persistence

- After test commands complete, the backend sends a summary block over SSE.
- Then it saves the result into Postgres via GORM.
- If save succeeds, the backend sends a final `event: done` message.
- The frontend listens for completion, closes the SSE stream, and refreshes the coverage history.

## Services configuration

Example service entry in `backend/config/services.json`:

```json
{
    "name": "hello-world",
    "display_name": "Hello World",
    "backend_path": "../../hello-world/backend",
    "frontend_path": "../../hello-world/frontend",
    "backend_type": "go",
    "frontend_type": "npm"
  }
```

- `name`: unique internal service key.
- `display_name`: friendly name shown in the dashboard.
- `backend_path`: relative backend test folder from /backend of coverage service.
- `frontend_path`:relative frontend test folder from /backend of coverage service.
- `backend_type`: currently supports `go`.
- `frontend_type`: currently supports `npm`.

## Database configuration

The backend reads Postgres connection values from environment variables:

- `DB_HOST`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_PORT`

If you use a `.env` file, it will be loaded automatically by `backend/main.go` and `backend/database/db.go`.

## Running locally

1. Start the backend:

```bash
cd coverage-service/backend
go run main.go
```

2. Start the frontend:

```bash
cd coverage-service/frontend
npm install
npm run dev
```

3. Open the dashboard:

```text
http://localhost:5173
```

4. Select a service and click **Run Tests**.

## API reference

| Method | Path                        | What it does                        |
|--------|-----------------------------|-------------------------------------|
| GET    | /api/registered-services    | Returns the services from `services.json` |
| GET    | /api/run/:service           | Runs tests and streams output via SSE |
| POST   | /api/coverage               | Saves a coverage result to the DB |
| GET    | /api/coverage               | Returns saved coverage history |
| GET    | /api/coverage/services      | Returns distinct service names saved in DB |
| GET    | /api/coverage/download      | Downloads CSV export of saved results |
