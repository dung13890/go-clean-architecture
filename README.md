# go-clean-architecture
![workflow status](https://github.com/dung13890/go-clean-architecture/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A modular, scalable Golang codebase built with Clean Architecture principles.

**Go version:** v1.25

## Overview

This project demonstrates clean architecture for Go applications with separation of concerns across distinct layers. Independent of frameworks, fully testable, and easy to extend.

📖 [Learn more about Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## 🚀 Quick Start

### 1. Setup

```bash
git clone https://github.com/dung13890/go-clean-architecture.git
cd go-clean-architecture
git submodule update --init --force --remote
```

### 2. Build & Run (Local)

```bash
docker compose build
cp .env.example .env
docker compose up -d
```

### 3. Inside Docker

```bash
docker compose exec go-app sh
```

### 4. Database Migration & Seed

```bash
make create_example_table.sql
go run cmd/migrate/main.go
# go run cmd/migrate/main.go down {step}
go run cmd/seed/main.go
```

### 5. Run App

```bash
air -c cmd/app/.air.toml
```

### 6. Useful Commands

```bash
make lint     # Check lint
make go-gen   # Generate mocks or base files
make test     # Run unit tests
```

## 🌐 API Test Example

```bash
curl -X POST 'localhost:8080/api/register' \
 -H 'accept: application/json' \
 -H 'content-type: application/json' \
 -d '{
    "email": "user@example.com",
    "password" : "password",
    "role_id": 1,
    "name": "user"
}'

curl -X POST 'localhost:8080/api/login' \
 -H 'accept: application/json' \
 -H 'content-type: application/json' \
 -d '{
    "email": "user@example.com",
    "password" : "password"
}'
```

## 🧱 Project Structure

```plaintext
Handler → UseCase → Domain Interface (Service, Repository)
                           ↑
                           └─ Adapters (Cache, DB, External)

```

```plaintext
internal/
├── domain/              # Core entities and interfaces
│   ├── entity/
│   ├── repository/
│   └── service/
├── usecase/             # Application logic
├── service/             # Business logic implementations
├── adapter/             # External integrations
│   ├── repository/      # Data persistence
│   ├── cache/           # Cache
│   └── external/        # External services
├── delivery/http/       # HTTP handlers
└── infrastructure/      # Config, database, logging
```

Go-App follows **Clean Architecture** with separation of concerns across distinct layers:

### 🗂 Layer Responsibilities

| Layer        | Path                | Description                                    |
| ------------ | ------------------- | ---------------------------------------------- |
| **Domain**   | `internal/domain`   | Core business entities and interfaces          |
| **Usecase**  | `internal/usecase`  | Application workflows and orchestration        |
| **Service**  | `internal/service`  | Shared reusable services (JWT, throttle, etc.) |
| **Adapter**  | `internal/adapter`  | External integrations (DB, cache, email)       |
| **Delivery** | `internal/delivery` | User interfaces (CLI, HTTP, Grpc future)       |
| **Registry** | `internal/registry` | Dependency injection and initialization        |


## ✨ Features

* **Authentication** — JWT-based authentication
* **User Management** — CRUD for users
* **Role Management** — Manage roles & permissions
* **Email Notifications** — SMTP email service
* **Rate Limiting** — Request throttling with Redis


## Stack

- [Echo](https://echo.labstack.com) — Web framework
- [GORM](https://gorm.io) — ORM
- [Viper](https://github.com/spf13/viper) — Configuration
- [JWT](https://golang-jwt.github.io/jwt) — Authentication
- [Air](https://github.com/cosmtrek/air) — Hot reload

---

## ☕ Support

If you find this project helpful:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dung13890)
