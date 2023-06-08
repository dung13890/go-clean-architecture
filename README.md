# go-clean-architecture
![workflow status](https://github.com/dung13890/go-clean-architecture/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Codebase for golang use clean architecture.

*The current go version use is `v1.20`*

## Overview
The purpose of the codebase is to show:
- Independent of Frameworks
- Testable
- Independent of UI
- Independent of Database
- Independent of any external agency


More at [https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Content
- [Quick start](#quick-start)
- [Project structure](#project-structure)
- [Tools Used](#tools-used)
- [Tool Generate](#tool-generate)

## Quick start
Below are some feature included in this project:
- auth (JWT / API)
- users (Create / List / Show / Update / Delete)
- roles (Create / List / Show / Update / Delete)

Build Local development:
```bash
cd go-clean-architecture
git submodule update --init --force --remote
docker compose build
```

Start development:
```bash
# Install dependencies
go mod tidy

# Copy env
cp .env.example .env

# Start docker
docker compose up -d

# Inside docker
docker compose exec go-app bash

# Make migrate
make create_example_table.sql

# Migrate
go run cmd/migrate/main.go

# Migrate down
# go run cmd/migrate/main.go down {step}
go run cmd/migrate/main.go down 2

# Run seed data
go run cmd/seed/main.go

# Start http server
air -c cmd/app/.air.toml

# Check lint
make lint

# Go generate mock, something
make go-gen

# Check Unit test
make test

# Check with CURL
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

## Tool Generate
If you need to generate a code base like this architecture, you can use [go-base-gen](https://github.com/dung13890/go-base-gen) tool.
You can read more about the tool at [README](https://github.com/dung13890/go-base-gen/blob/master/README.md#usage)
```bash
NAME:
   go-base-gen - Use this tool to generate base code

USAGE:
   go-base-gen [global options] command [command options] [arguments...]

VERSION:
   v1.0.10

COMMANDS:
   project  Generate base code for go project use clean architecture
   domain   Create new domain in project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print only the version (default: false)
```

The [go-base-gen](https://github.com/dung13890/go-base-gen) tool was created with the purpose to help developers save their time in creating a project with a clean architecture. The tool will generate the code base for you, you just need to focus on the business logic.

## Project structure
![Clean Architecture](CleanArchitecture.jpeg)

This application is divided into 2 layers, internal and pkg:
- Internal is Business logic
- Pkg is tools (logs, database, utils,...)

The communication between layers

![Clean Architecture Layer](clean_layer.svg)

```mermaid
flowchart LR
    ex[External]
    de[Delivery]
    uc[Usecase]
    rp[Repository]
    ps[Pubsub]
    ot[Other]
    db[(Database)]

    subgraph in [Internal]
        direction TB
        de -.->|Interface / Domain| uc -.->|Interface / Domain| ps & ot & rp
    end

    ex -.->|DTO|in -.->|DAO|db

```


For Internal application use 4 layers:

### `Entities | domain`:
Entities / domain is the most inner layer of the onion architecture. It is a struct for data that will be used by communication between layers.

Entities are simple data structures:
```go
// Path internal/domain/role.go
// Role entity
type Role struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Slug      string    `json:"slug"`
    CreatedAt time.Time `json:"created_at"`
}
```

### `Repositories`:
A repository is an abstract storage (database) that business logic works with. Layer responsibility will choose DB use in application
```go
// RoleRepository represent the role's usecases
type RoleRepository interface {
    Fetch(context.Context) ([]Role, error)
}
```

### `Usecase`:
This layer contains application specific business rules. This a layer decide repository, service, other use in application.

```go
// RoleUsecase represent the role's repository contract
type RoleUsecase interface {
    Fetch(context.Context) ([]Role, error)
}
```

### `Delivery`:
This a layer will decide how the data present. Could be REST API, HTML, or gRPC whatever the decide type.
```go
// Path: internal/modules/role/delivery/http

// roleHandler represent the httphandler
type roleHandler struct {
    Usecase domain.RoleUsecase
}

// NewHandler will initialize the roles/ resources endpoint
func NewHandler(e *echo.Echo, uc domain.RoleUsecase) {
    handler := &roleHandler{
        Usecase: uc,
    }

    g := e.Group("/api")
    g.GET("/roles", handler.Index)
}

// Index will fetch data
func (hl *roleHandler) Index(c echo.Context) error {
    ctx := c.Request().Context()
    roles, _ := hl.Usecase.Fetch(ctx)

    return c.JSON(http.StatusOK, roles)
}
```

## Tools Used
- [https://gorm.io](https://gorm.io)
- [validator](https://github.com/go-playground/validator)
- [spf13/viper](https://github.com/spf13/viper)
- [golang/mock](https://github.com/golang/mock)
- [Echo](https://echo.labstack.com)
- [JWT](https://golang-jwt.github.io/jwt)
- [cosmtrek/air](https://github.com/cosmtrek/air)
- [go-base-gen](https://github.com/dung13890/go-base-gen)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dung13890)
