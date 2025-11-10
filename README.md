# go-clean-architecture
![workflow status](https://github.com/dung13890/go-clean-architecture/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A modular, scalable, and testable **Golang codebase** built using **Clean Architecture** principles.

> Current Go version: **v1.25**

## 🧩 Overview

This project demonstrates a clean and maintainable architecture for Go applications:

* Independent of frameworks and external layers
* Fully testable
* Database and UI agnostic
* Easy to extend with new modules

📖 Learn more about the Clean Architecture:
[Uncle Bob’s Blog – The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Content
- [Quick start](#quick-start)
- [Project structure](#project-structure)
- [Feature](#features)
- [Tools & Dependencies](#tools--dependencies)

---

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
go mod tidy
cp .env.example .env
docker compose up -d
```

### 3. Inside Docker

```bash
docker compose exec go-app bash
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

---

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

---

## 🧱 Project Structure

```plaintext
Handler
   ↓
UseCase ──────→ Domain Interface (Service, Repository)
   ↑                    ↑
   │provider            │ implements
   │                    │
Registry ──→ Factory ───┴─→ Adapters (Cache, External Service, Database)
   ↑
   │
 Config

```

Go-App follows **Clean Architecture** with separation of concerns across distinct layers:

### 🗂 Directory Overview

| Layer        | Path                | Description                                    |
| ------------ | ------------------- | ---------------------------------------------- |
| **Domain**   | `internal/domain`   | Core business entities and interfaces          |
| **Usecase**  | `internal/usecase`  | Application logic and business rules           |
| **Service**  | `internal/service`  | Shared reusable services (JWT, throttle, etc.) |
| **Adapter**  | `internal/adapter`  | External systems (DB, cache, email)            |
| **Registry** | `internal/registry` | Dependency injection and initialization        |


## ✨ Features

* **Authentication** — JWT-based authentication
* **User Management** — CRUD for users
* **Role Management** — Manage roles & permissions
* **Email Notifications** — SMTP email service
* **Rate Limiting** — Request throttling with Redis

---

## 🛠 Tools & Dependencies

* [GORM](https://gorm.io) — ORM for database interaction
* [Echo](https://echo.labstack.com) — HTTP web framework
* [Viper](https://github.com/spf13/viper) — Configuration management
* [Validator](https://github.com/go-playground/validator) — Input validation
* [Golang/mock](https://github.com/golang/mock) — Mock generation
* [JWT](https://golang-jwt.github.io/jwt) — Token authentication
* [cosmtrek/air](https://github.com/cosmtrek/air) — Hot reload
* [go-base-gen](https://github.com/dung13890/go-base-gen) — Code generation tool

---

## ☕ Support

If you find this project helpful:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dung13890)

---
