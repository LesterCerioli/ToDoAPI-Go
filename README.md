
# To Do Golang Rest API


A production-ready REST API for managing todo tasks, built with **Go 1.26** and **Fiber v3**. Features SQLite persistence, interactive Swagger documentation, graceful shutdown, and Docker support.

[![Go Version](https://img.shields.io/badge/Go-1.26-blue.svg)](https://golang.org/)
[![Fiber Version](https://img.shields.io/badge/Fiber-v3-00ADD8.svg)](https://gofiber.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## 🚀 Live Demo

**Base URL:** `https://todoapi-go.onrender.com`

| Resource | URL |
|----------|-----|
| API Base | `https://todoapi-go.onrender.com/api/v1` |
| Health Check | `https://todoapi-go.onrender.com/health` |
| Swagger UI | `https://todoapi-go.onrender.com/swagger/` |
| Swagger JSON | `https://todoapi-go.onrender.com/swagger/doc.json` |

---

## 📦 Tech Stack

| Technology | Purpose |
|------------|---------|
| [Go 1.26](https://golang.org/) | Programming language |
| [Fiber v3](https://gofiber.io/) | High-performance HTTP web framework |
| [GORM](https://gorm.io/) | ORM for database operations |
| [SQLite](https://sqlite.org/) | Lightweight embedded database |
| [Swagger UI](https://swagger.io/tools/swagger-ui/) | Interactive API documentation |
| [Docker](https://www.docker.com/) | Containerization |

---

## 🔧 Getting Started

### Prerequisites

- Go 1.26+
- (Optional) Docker & Docker Compose
- (Optional) [swag](https://github.com/swaggo/swag) CLI for doc generation

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/todo-api.git
cd todo-api

# Install dependencies
go mod tidy

# Generate Swagger docs (requires swag CLI)
swag init

# Run the application
go run main.go
```

The server starts at `http://localhost:3000`.

---

## 📚 API Endpoints

### Base Path: `/api/v1`

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| `GET` | `/health` | Health check & system status | No |
| `POST` | `/todos` | Create a new todo | No |
| `GET` | `/todos` | List all todos | No |
| `GET` | `/todos/:id` | Get a specific todo by ID | No |
| `PATCH` | `/todos/:id` | Update a todo (partial) | No |
| `DELETE` | `/todos/:id` | Delete a todo | No |

### Request/Response Examples

#### Create Todo
```bash
curl -X POST https://todoapi-go.onrender.com/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Fiber v3",
    "description": "Build a production-ready API"
  }'
```

**Response (201 Created):**
```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Learn Fiber v3",
    "description": "Build a production-ready API",
    "completed": false,
    "created_at": "2026-04-25T12:00:00Z",
    "updated_at": "2026-04-25T12:00:00Z"
  }
}
```

#### List All Todos
```bash
curl https://todoapi-go.onrender.com/api/v1/todos
```

**Response (200 OK):**
```json
{
  "count": 2,
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Learn Fiber v3",
      "description": "Build a production-ready API",
      "completed": false,
      "created_at": "2026-04-25T12:00:00Z",
      "updated_at": "2026-04-25T12:00:00Z"
    }
  ]
}
```

#### Update Todo (Partial)
```bash
curl -X PATCH https://todoapi-go.onrender.com/api/v1/todos/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

#### Delete Todo
```bash
curl -X DELETE https://todoapi-go.onrender.com/api/v1/todos/550e8400-e29b-41d4-a716-446655440000
```

**Response:** `204 No Content`

---

## 🐳 Docker

### Build & Run

```bash
# Using Docker Compose
docker-compose up --build

# Or using Docker directly
docker build -t todo-api .
docker run -p 3000:3000 -v todo-data:/data todo-api
```

### Docker Compose Services

| Service | Port | Volume |
|---------|------|--------|
| `api` | `3000:3000` | `todo-data:/data` (SQLite persistence) |

---

## 🏗️ Project Structure

```
todo-api/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── Dockerfile              # Container image
├── docker-compose.yml      # Multi-container orchestration
├── README.md               # This file
├── database/
│   └── db.go               # Database connection & migration
├── docs/
│   ├── docs.go             # Swagger generated docs
│   ├── swagger.json        # OpenAPI 2.0 JSON spec
│   └── swagger.yaml        # OpenAPI 2.0 YAML spec
├── handlers/
│   └── todo.go             # HTTP request handlers
├── middleware/
│   └── logger.go           # Fiber logger middleware
├── model/
│   └── todo.go             # Data models & request DTOs
└── storage/
    └── sqlite.go           # SQLite repository implementation
```

---

## 🧪 Testing with Swagger UI

Visit `https://todoapi-go.onrender.com/swagger/` to explore and test all endpoints interactively:

1. Click any endpoint to expand it
2. Click **"Try it out"**
3. Fill in parameters/request body
4. Click **"Execute"**
5. View the server response directly in the browser

---

## ⚙️ Configuration

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `PORT` | `3000` | HTTP server port |
| `DB_PATH` | `todos.db` | SQLite database file path |

---

## 🛡️ Graceful Shutdown

The server handles `SIGINT` and `SIGTERM` signals with a **10-second timeout** for active requests to complete before shutting down.

---

## 📄 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Your Name** — [@yourusername](https://github.com/yourusername)

---

<p align="center">
  Built with ❤️ using <a href="https://golang.org/">Go</a> & <a href="https://gofiber.io/">Fiber</a>
</p>
```

---

## Quick Customization Checklist

Replace these placeholders before committing:

| Placeholder | Replace With |
|-------------|--------------|
| `yourusername` | Your actual GitHub username |
| `yourusername/todo-api` | Your actual repo path |
| `support@example.com` | Your contact email (or remove) |
| `Your Name` | Your actual name |

---

