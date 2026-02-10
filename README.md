# 🗂️ Task Management System – Go Backend

A production-ready Task Management REST API built with Golang, Gin, PostgreSQL, JWT, GORM, and Docker.

---

## 🚀 Features

- User Authentication (JWT)
- Role-Based Access (Admin/User)
- Project Management (CRUD)
- Task Management (CRUD + Workflow)
- Admin Panel
- Logging, Recovery, Rate Limiting

---

## 🛠️ Tech Stack

- Go (Golang)
- Gin Framework
- PostgreSQL
- GORM
- JWT
- Docker

---

## 📁 Project Structure

task-manager/
├── cmd/
├── config/
├── internal/
├── routes/
├── pkg/
├── docker/
└── README.md

---

## ⚙️ Environment Setup

Create .env file:

DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=task_manager
DB_PORT=5432
JWT_SECRET=your_secret_key

---

## ▶️ Run Locally

go mod tidy
go run cmd/main.go

---

## 🔐 Authentication

Use JWT in headers:

Authorization: Bearer <TOKEN>

---

## 📡 API Endpoints

Auth:
POST /api/auth/register
POST /api/auth/login

Projects:
POST /api/projects
GET /api/projects
GET /api/projects/:id
PUT /api/projects/:id
DELETE /api/projects/:id

Tasks:
POST /api/tasks
GET /api/tasks
GET /api/tasks/:id
PUT /api/tasks/:id
DELETE /api/tasks/:id

Admin:
GET /api/admin/users
PUT /api/admin/users/:id/role

---

## 📊 Database Tables

users, projects, tasks

---

## 💼 Resume Line

Built scalable Task Manager API using Go, Gin, PostgreSQL, JWT, and Docker.

---

## 👨‍💻 Author

Anurag Dahiwade
GitHub: https://github.com/AnuragDahiwade

---

## 📜 License

MIT License
