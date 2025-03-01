# 📝 Task Management System

## 📌 Overview
This is a **Task Management Microservice** built using **Go (Fiber framework)** with **PostgreSQL** database.  

## 🚀 Features
- **RESTful API** with clear endpoints.  
- **Create, Read, Update, Delete (CRUD) operations** for tasks.  
- **Pagination & Filtering** for listing tasks.  
- **Database Migrations** using GORM.  
- **Dockerized Setup** for easy deployment.  
- **Adminer** UI for database management.  

---

## 📂 Folder Structure
```
├── README.md
├── cmd
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── handlers
│   │   ├── routes.go
│   │   └── task_handler.go
│   ├── models
│   │   └── task.go
│   ├── repository
│   │   └── task_repository.go
│   └── services
│       └── task_service.go
├── migrations
│   ├── 000001_create_tasks_table.down.sql
│   └── 000001_create_tasks_table.up.sql
└── pkg
    ├── config
    │   └── config.go
    └── db
        ├── db.go
        └── migrate.go

```


---

## ⚡ Installation & Setup

### **Clone the Repository**
```sh
git clone https://github.com/your-username/task-service.git
cd task-service
```

### Install Dependencies

```sh
go mod tidy
```

### Configure Environment Variables

Create a .env file in the root directory:

```sh
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=tasks_db
DB_PORT=5432
SERVER_PORT=8080
```

### Run Database With Docker

```sh
docker-compose up -d
```

### Run Migrations

```sh
go run cmd/main.go migrate
```

### Run Server

```sh
go run cmd/main.go
```
Run the server on http://localhost:8080

### Access Adminer UI

```sh
http://localhost:8081
```

---

##  API Endpoints

### Create a Task

- `POST /api/tasks`
- Request Body:

```json
{
  "title": "Build API",
  "description": "Implement a Task Management API using Go and GORM.",
  "status": "Pending"
}

```

- Response:

```json
{
  "message": "Task created successfully",
  "task": {
    "id": 1,
    "title": "Build API",
    "description": "Implement a Task Management API using Go and GORM.",
    "status": "Pending",
    "created_at": "2025-03-01T10:00:00Z",
    "updated_at": "2025-03-01T10:00:00Z"
  }
}
```

### Get All Tasks

- `GET /api/tasks`
- Query Parameters:
    - `limit`: Maximum number of tasks to return (default: 10).
    - `offset`: Number of tasks to skip (default: 0).
    - `status`: Filter tasks by status (default: all statuses).
- Response:

```json
{
  "tasks": [
    {
      "id": 11,
      "title": "Task 11",
      "status": "Pending",
      "created_at": "2025-03-01T10:00:00Z",
      "updated_at": "2025-03-01T10:00:00Z"
    }
  ],
  "total_pages": 4,
  "current_page": 3,
  "has_next": true
}

```

### Get a Single Task

- `GET /api/tasks/{id}`
- Path Parameter:
    - `id`: ID of the task to retrieve.
- Response:

```json
{
    "id": 2,
    "title": "Test2",
    "description": "Testing description",
    "status": "Pending",
    "created_at": "2025-03-01T20:07:39.303385+05:30",
    "updated_at": "2025-03-01T20:07:39.303385+05:30"
}
```

### Update a Task

- `PUT /api/tasks/{id}`
- Path Parameter:
    - `id`: ID of the task to update.
- Request Body:

```json
{
  "title": "Updated Task",
  "status": "InProgress"
}

```

### Delete a Task

- `DELETE /api/tasks/{id}`
- Path Parameter:
    - `id`: ID of the task to delete.
- Response:

```json
{
  "message": "Task deleted successfully"
}

```