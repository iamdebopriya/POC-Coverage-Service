# Inventory Tracker

A modern, full-stack web application for managing inventory items with real-time tracking, user authentication, and an intuitive interface. Built with Vue 3 and Go.

---

## Features

- **User Authentication** — Register and login with secure password handling
- **Inventory Management** — Create, read, update, and delete inventory items
- **Real-time Tracking** — Monitor stock levels and item details
- **Low Stock Alerts** — Get notified when items fall below minimum thresholds
- **Responsive Design** — Works seamlessly on desktop and mobile devices
- **RESTful API** — Clean, well-documented backend APIs

---

## Tech Stack

| Layer | Technology |
|---|---|
| Frontend | Vue 3 + Vite |
| Backend | Go + Gin Web Framework |
| Database | PostgreSQL |
| ORM | GORM |

---

## Project Structure

```
inventory-tracker/
├── backend/           # Go backend API
│   ├── main.go       # Application entry point
│   ├── database/     # Database connection and setup
│   ├── handlers/     # API request handlers
│   ├── models/       # Data models
│   └── router/       # API routes
├── frontend/          # Vue 3 frontend application
│   ├── src/
│   │   ├── pages/    # Page components
│   │   ├── components/ # Reusable components
│   │   ├── services/ # API client
│   │   └── router/   # Vue Router configuration
│   ├── index.html
│   ├── package.json
│   └── vite.config.js
└── README.md
```

---

## Prerequisites

- **Go** 1.21 or higher
- **Node.js** v18+ and npm
- **PostgreSQL** 16 or higher
- **Git**

---

## Getting Started

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Set up environment variables in `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=inventory_tracker
PORT=8080
```

3. Install dependencies and run the server:
```bash
go mod download
go run main.go
```

The API will be available at `http://localhost:8080`

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Configure the API endpoint in `.env`:
```env
VITE_API_BASE_URL=http://localhost:8080/api
```

3. Install dependencies and start the development server:
```bash
npm install
npm run dev
```

The application will be available at `http://localhost:5173`

---

## Available Scripts

### Backend
- `go run main.go` — Start the development server
- `go build` — Build the binary for production

### Frontend
- `npm run dev` — Start development server
- `npm run build` — Build for production
- `npm run preview` — Preview production build locally

---

## API Endpoints

### Authentication
- `POST /api/auth/register` — Register a new user
- `POST /api/auth/login` — Login user

### Items
- `GET /api/items` — Get all inventory items
- `POST /api/items` — Create a new item
- `GET /api/items/:id` — Get item details
- `PUT /api/items/:id` — Update an item
- `DELETE /api/items/:id` — Delete an item
- `GET /api/items/low-stock` — Get low stock items

---

## Usage

1. Open the application at `http://localhost:5173`
2. Create a new account or login
3. Add inventory items with details like name, quantity, and minimum threshold
4. View and manage your inventory in real-time
5. Receive alerts when items fall below minimum levels

---
