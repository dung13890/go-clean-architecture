# go-clean-architecture

![workflow status](https://github.com/dung13890/go-clean-architecture/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A modular, scalable Go codebase template built with Clean Architecture principles.

**Go version:** v1.25

## 🎯 What Is This?

This is a **production-ready template** for building Go applications with:
- ✅ Clean separation of concerns
- ✅ Framework-independent business logic
- ✅ Fully testable code
- ✅ Easy to extend and maintain

📖 [Learn more about Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

---

## 🚀 Two Ways to Get Started

### Option 1: Use the Generator Tool (Recommended)

The easiest way to create a new project based on this architecture:

```bash
# Install the generator
go install github.com/dung13890/go-base-gen@latest

# Generate your new project
go-base-gen project --pkg github.com/yourusername/yourproject --path yourproject

# Setup and run
cd yourproject
go mod tidy
cp .env.example .env
go run cmd/migrate/main.go
make dev
```

👉 [View go-base-gen documentation](https://github.com/dung13890/go-base-gen)

### Option 2: Clone This Repository

To explore or contribute to the template itself:

```bash
# Clone the repository
git clone https://github.com/dung13890/go-clean-architecture.git
cd go-clean-architecture
git submodule update --init --force --remote

# Build and run with Docker
docker compose build
cp .env.example .env
docker compose up -d

# Access the container
docker compose exec go-app sh

# Inside container: Setup database
make create_example_table.sql
go run cmd/migrate/main.go
go run cmd/seed/main.go

# Run the application
air -c cmd/app/.air.toml
```

---

## 📁 Project Structure

### Architecture Flow

```plaintext
Handler → UseCase → Domain Interface (Service, Repository)
    ↓                       ↑
Validation                  └─ Adapters (Cache, DB, External)
```

### Directory Layout

```plaintext
internal/
├── domain/              # 🎯 Core Business Layer
│   ├── entity/          # Business entities (User, Role, etc.)
│   ├── repository/      # Repository interfaces
│   └── service/         # Service interfaces
│
├── usecase/             # 📋 Application Logic
├── service/             # 🔧 Business Logic Implementation
├── adapter/             # 🔌 External Integrations
│   ├── repository/      # Data persistence
│   ├── cache/           # Cache
│   └── external/        # Third-party APIs
│
├── delivery/            # 🌐 User Interface Layer
│   └── http/            # HTTP handlers & routes
├── infrastructure/      # ⚙️ Cross-cutting Concerns
│
└── registry/            # 🏗️ Dependency Injection
```

### Layer Responsibilities

| Layer            | Responsibility                          | Example                    |
| ---------------- | --------------------------------------- | -------------------------- |
| **Domain**       | Core business rules & interfaces        | User entity, UserRepo      |
| **Usecase**      | Application workflows & orchestration   | Register user flow         |
| **Service**      | Reusable business logic                 | JWT generation, Email send |
| **Adapter**      | External system integration             | PostgreSQL, Redis, SMTP    |
| **Delivery**     | Request/response handling               | HTTP handlers, CLI         |
| **Registry**     | Wire dependencies together              | Inject repos into usecases |
| **Infrastructure** | Technical concerns                    | Config, DB, Redis          |

---

## ✨ Built-in Features

- 🔐 **JWT Authentication** — Secure token-based auth
- 👤 **User Management** — Full CRUD operations
- 🎭 **Role & Permissions** — Access control system
- 📧 **Email Service** — SMTP integration
- 🚦 **Rate Limiting** — Redis-based throttling
- 🔄 **Hot Reload** — Development with Air
- 🧪 **Testing Ready** — Mock generation included

---

## 🛠️ Development Commands

```bash
# Code quality
make lint              # Run linter
make test              # Run unit tests
make go-gen            # Generate mocks

# Database
make create_example_table.sql  # Create migration file
go run cmd/migrate/main.go     # Run migrations
go run cmd/migrate/main.go down 1  # Rollback 1 step
go run cmd/seed/main.go        # Seed database

# Development
make dev               # Run with hot reload
```

---

## 🧪 API Examples

### Register User

```bash
curl -X POST 'http://localhost:8080/api/register' \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "user@example.com",
    "password": "password",
    "role_id": 1,
    "name": "John Doe"
  }'
```

### Login

```bash
curl -X POST 'http://localhost:8080/api/login' \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "user@example.com",
    "password": "password"
  }'
```

---

## 📦 Tech Stack

| Category       | Technology                                             |
| -------------- | ------------------------------------------------------ |
| Framework      | [Echo](https://echo.labstack.com)                      |
| ORM            | [GORM](https://gorm.io)                                |
| Configuration  | [Viper](https://github.com/spf13/viper)                |
| Authentication | [JWT](https://golang-jwt.github.io/jwt)                |
| Hot Reload     | [Air](https://github.com/cosmtrek/air)                 |
| Database       | PostgreSQL                                             |
| Cache          | Redis                                                  |

---

## 🎓 Learn More

### Adding New Features

**Want to add a new domain (e.g., "Product")?**

```bash
# Use the generator tool
go-base-gen domain --dn product --pkg github.com/yourusername/yourproject

# This creates:
# - internal/domain/entity/product.go
# - internal/domain/repository/product_repo.go
# - internal/domain/service/product_svc.go
# - internal/service/product_svc.go
# - internal/usecase/product/
# - internal/adapter/repository/product_dao.go
# - internal/adapter/repository/product_repo.go
# - internal/delivery/http/product_handler.go
# - internal/delivery/http/dto/product_dto.go
```

### Architecture Principles

1. **Dependency Rule**: Inner layers don't know about outer layers
2. **Interface Segregation**: Use small, focused interfaces
3. **Dependency Injection**: Wire dependencies in registry
4. **Testability**: Mock external dependencies

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

## ☕ Support

If you find this project helpful:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dung13890)

---

## 📄 License

MIT License - see [LICENSE](https://opensource.org/licenses/MIT) for details.
