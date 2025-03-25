# Go Backend for Young Conf

A scalable backend built with **Go**, **Gin**, **GORM**, **PostgreSQL**, and **Redis**.

## 🚀 Features

- **RESTful API** using Gin
- **PostgreSQL + Redis** integration
- **Modular service-based architecture**
- **Middleware support** (CORS, Logging)
- **Graceful shutdown handling**
- **Environment variable configuration**

---

## 📌 Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/doc/install) (1.18+ recommended)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/download/)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/hokamsingh/go-backend-template.git
   cd go-backend-template
   ```

2. Create a `.env` file and configure database credentials:

   ```bash
   cp .env.example .env
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the database migrations (if applicable):

   ```bash
   go run scripts/migrate.go
   ```

5. Start the server:

   ```bash
   go run main.go
   ```

   The server should be running at **http://localhost:8080**

---

## 📂 Project Structure

```
├── config/          # Configuration files
├── controllers/     # API Controllers
├── models/         # Database Models
├── routes/         # Route Handlers
├── services/       # Business Logic Layer
├── database/       # DB Connection & Migrations
├── middleware/     # Middleware (CORS, Auth, Logging)
├── templates/      # HTML Templates
├── main.go         # Entry Point
└── .env.example    # Environment Config Sample
```

---

## 📡 API Endpoints

| Method | Endpoint         | Description                |
|--------|------------------|----------------------------|
| GET    | `/api/events`    | Get all events             |
| GET    | `/api/speakers`  | Create all speakers        |
| GET    | `/api/speakers/?id={id}` | Get speaker by ID  |

---

## 🛠 Technologies Used

- **Go** - Core language
- **Gin** - HTTP framework
- **GORM** - ORM for PostgreSQL
- **Redis** - In-memory caching
- **Docker** - Containerization (optional)

---