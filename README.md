# Go API Builder

A modern, production-ready REST API built with Go using clean architecture principles, SQLC for type-safe database queries, PostgreSQL, Redis, and Docker containerization.

## 🚀 Project Overview

This project demonstrates a well-structured Go REST API following industry best practices. It features a clean layered architecture with separation of concerns, type-safe database operations using SQLC, comprehensive error handling, and production-ready deployment configurations.

### 🎯 Key Features

- **Clean Architecture**: Repository pattern with service layers
- **Type-Safe Database Access**: SQLC-generated queries with full type safety
- **Modern Stack**: Go 1.24, Gin framework, PostgreSQL, Redis
- **Container-First**: Docker and Docker Compose for development and production
- **Security**: PBKDF2 password hashing with secure salt generation
- **Health Checks**: Built-in health monitoring and graceful shutdown
- **Hot Reload**: Air for development with live reloading
- **Database Migrations**: Structured migration system
- **Environment Configuration**: Flexible configuration management

## 📋 Table of Contents

- [Architecture](#-architecture)
- [Tech Stack](#-tech-stack)
- [Getting Started](#-getting-started)
- [API Documentation](#-api-documentation)
- [Database Schema](#-database-schema)
- [Project Structure](#-project-structure)
- [ORM Comparison](#-orm-comparison-sqlc-vs-ent-vs-gorm)
- [Development](#-development)
- [Deployment](#-deployment)
- [Contributing](#-contributing)

## 🏗️ Architecture

The project follows **Clean Architecture** principles with clear separation of concerns:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Presentation  │    │    Business     │    │      Data       │
│     Layer       │────│     Logic       │────│     Access      │
│   (Handlers)    │    │   (Services)    │    │ (Repositories)  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
    ┌────▼────┐             ┌────▼────┐             ┌────▼────┐
    │   Gin   │             │ Business│             │  SQLC   │
    │ Router  │             │  Rules  │             │Generated│
    └─────────┘             └─────────┘             └─────────┘
```

### Layer Responsibilities

1. **Handler Layer**: HTTP request/response handling, input validation, response formatting
2. **Service Layer**: Business logic, data validation, transaction management
3. **Repository Layer**: Database operations, data access abstraction
4. **Database Layer**: SQLC-generated type-safe queries

## 🛠️ Tech Stack

### Core Technologies
- **Language**: Go 1.24.1
- **Web Framework**: Gin (high-performance HTTP web framework)
- **Database**: PostgreSQL 16 with pgx/v5 driver
- **Cache**: Redis 7
- **Query Builder**: SQLC (compile-time SQL query generation)

### Development & DevOps
- **Containerization**: Docker & Docker Compose
- **Hot Reload**: Air for development
- **Migration**: golang-migrate
- **Security**: PBKDF2 password hashing with crypto/rand

### Dependencies
```go
// Core Dependencies
github.com/gin-gonic/gin v1.10.1        // Web framework
github.com/jackc/pgx/v5 v5.7.5          // PostgreSQL driver
github.com/go-redis/redis/v8 v8.11.5     // Redis client
golang.org/x/crypto v0.42.0             // Cryptographic functions
```

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development)
- Make (optional, for convenience commands)

### Quick Start with Docker

1. **Clone the repository**
   ```bash
   git clone https://github.com/veliulugut/go-apibuilder.git
   cd go-apibuilder
   ```

2. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your preferred settings
   ```

3. **Start the application**
   ```bash
   # Development mode with hot reload
   docker-compose -f compose.yml -f compose.dev.yml up --build

   # Production mode
   docker-compose up --build
   ```

4. **Verify the setup**
   ```bash
   curl http://localhost:8080/ping
   # Response: {"message":"pong"}
   ```

### Local Development Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Install development tools**
   ```bash
   go install github.com/air-verse/air@latest
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

3. **Start services**
   ```bash
   # Start PostgreSQL and Redis
   docker-compose up db redis -d

   # Run migrations
   migrate -path db/migration -database "postgres://user:password@localhost:5432/mydatabase?sslmode=disable" up

   # Start the application with hot reload
   air
   ```

## 📚 API Documentation

### Base URL
```
Development: http://localhost:8080
API Prefix: /api/v1
```

### Endpoints

#### Health Check
```http
GET /ping
```
**Response:**
```json
{
  "message": "pong"
}
```

#### User Management

##### Create User
```http
POST /api/v1/users
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "password": "securepassword123"
}
```

**Success Response (201):**
```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

**Error Response (400):**
```json
{
  "error": "Invalid request payload: Email is required"
}
```

##### Get User by ID
```http
GET /api/v1/users/{id}
```

**Success Response (200):**
```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

**Error Response (404):**
```json
{
  "error": "User not found"
}
```

## 🗃️ Database Schema

### Users Table
```sql
CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### Migration Commands
```bash
# Create new migration
migrate create -ext sql -dir db/migration -seq migration_name

# Run migrations
migrate -path db/migration -database $DATABASE_URL up

# Rollback migrations
migrate -path db/migration -database $DATABASE_URL down 1
```

## 📁 Project Structure

```
go-apibuilder/
├── cmd/server/main.go              # Application entry point
├── config/config.go                # Configuration management
├── db/
│   ├── migration/                  # Database migrations
│   │   ├── 000001_create_users_table.up.sql
│   │   └── 000001_create_users_table.down.sql
│   ├── queries/user.sql            # SQL queries for SQLC
│   └── sqlc/                       # SQLC generated code
│       ├── db.go                   # Database connection interface
│       ├── models.go               # Generated models
│       ├── querier.go              # Generated query interface
│       └── user.sql.go             # Generated user queries
├── internal/
│   ├── handler/user.go             # HTTP handlers
│   ├── repository/user.go          # Data access layer
│   ├── router/user.go              # Route definitions
│   ├── service/user.go             # Business logic layer
│   └── util/password.go            # Utility functions
├── scripts/                        # Utility scripts
├── docker-compose.yml              # Production compose
├── docker-compose.dev.yml          # Development compose
├── Dockerfile                      # Multi-stage build
├── sqlc.yaml                       # SQLC configuration
├── go.mod                          # Go module definition
└── README.md                       # Project documentation
```

### Layer Descriptions

- **`cmd/`**: Application entry points and main functions
- **`config/`**: Configuration loading and management
- **`db/`**: Database-related files (migrations, queries, generated code)
- **`internal/handler/`**: HTTP request handlers (Gin controllers)
- **`internal/service/`**: Business logic and use cases
- **`internal/repository/`**: Data access layer and database operations
- **`internal/router/`**: HTTP route definitions and middleware
- **`internal/util/`**: Shared utility functions and helpers

## 🔍 ORM Comparison: SQLC vs Ent vs GORM

### Why SQLC?

This project uses **SQLC** instead of traditional ORMs like GORM or Ent. Here's a comprehensive comparison:

#### SQLC ✅

**Strengths:**
- ✅ **Compile-time Safety**: SQL queries are validated at compile time
- ✅ **Zero Runtime Overhead**: Generates plain Go code, no reflection
- ✅ **SQL-First Approach**: Write actual SQL, get type-safe Go code
- ✅ **Performance**: No ORM overhead, direct SQL execution
- ✅ **Transparency**: Generated code is readable and debuggable
- ✅ **Migration Friendly**: Works seamlessly with migration tools
- ✅ **Learning Curve**: Developers learn SQL properly

**Use Case:** Perfect for performance-critical applications where you want full control over SQL queries.

```go
// SQLC Example - Type-safe and performant
user, err := queries.GetUserByID(ctx, userID)
if err != nil {
    return nil, err
}
```

#### GORM ⚖️

**Strengths:**
- ✅ **Rapid Development**: Quick CRUD operations
- ✅ **Rich Ecosystem**: Many plugins and extensions
- ✅ **Familiar ORM Patterns**: Similar to ORMs in other languages
- ✅ **Automatic Migrations**: Schema management built-in

**Weaknesses:**
- ❌ **Runtime Overhead**: Heavy use of reflection
- ❌ **Complex Query Generation**: Difficult to optimize complex queries
- ❌ **Hidden SQL**: Hard to debug generated SQL
- ❌ **N+1 Query Problems**: Easy to write inefficient code

```go
// GORM Example - Convenient but less transparent
var user User
result := db.First(&user, userID)
if result.Error != nil {
    return nil, result.Error
}
```

#### Ent ⚖️

**Strengths:**
- ✅ **Code Generation**: Type-safe code generation like SQLC
- ✅ **Schema-First**: Clear schema definitions
- ✅ **Rich Type System**: Advanced Go type support
- ✅ **Graph Queries**: Good for complex relationships

**Weaknesses:**
- ❌ **Learning Curve**: Steep learning curve and complex API
- ❌ **Vendor Lock-in**: Heavily tied to Ent's way of doing things
- ❌ **Query Complexity**: Complex queries can be verbose
- ❌ **Debugging**: Generated code can be hard to debug

```go
// Ent Example - Powerful but complex
user, err := client.User.
    Query().
    Where(user.ID(userID)).
    Only(ctx)
if err != nil {
    return nil, err
}
```

### Comparison Table

| Feature | SQLC | GORM | Ent |
|---------|------|------|-----|
| **Performance** | 🔥 Excellent | ⚡ Good | ⚡ Good |
| **Type Safety** | 🔥 Compile-time | ⚠️ Runtime | 🔥 Compile-time |
| **Learning Curve** | 🟢 Easy | 🟢 Easy | 🔴 Steep |
| **SQL Control** | 🔥 Full | ⚠️ Limited | ⚠️ Limited |
| **Development Speed** | ⚡ Fast | 🔥 Very Fast | ⚡ Fast |
| **Debugging** | 🔥 Excellent | ⚠️ Difficult | ⚠️ Difficult |
| **Migration Support** | 🔥 Excellent | 🟢 Good | 🟢 Good |
| **Ecosystem** | 🟢 Growing | 🔥 Mature | ⚡ Developing |

### Our Choice: SQLC

We chose SQLC for this project because:

1. **Performance First**: Zero runtime overhead means better performance
2. **SQL Mastery**: Team can write optimized SQL queries
3. **Debugging**: Easy to debug and optimize generated Go code
4. **Compile Safety**: Catch SQL errors at compile time, not runtime
5. **Simplicity**: Simple, predictable behavior without magic
6. **PostgreSQL Optimization**: Can leverage PostgreSQL-specific features

**Perfect for:** APIs requiring high performance, teams comfortable with SQL, projects needing query optimization.

## 🔧 Development

### Code Generation

```bash
# Generate SQLC code
sqlc generate

# Add new query to db/queries/user.sql then regenerate
```

### Database Operations

```bash
# Create new migration
migrate create -ext sql -dir db/migration -seq add_new_table

# Apply migrations
migrate -path db/migration -database $DATABASE_URL up

# Rollback last migration
migrate -path db/migration -database $DATABASE_URL down 1
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestCreateUser ./internal/service
```

### Development Workflow

1. **Database Changes**:
   - Create migration files
   - Update SQL queries in `db/queries/`
   - Run `sqlc generate`

2. **API Changes**:
   - Update handlers for new endpoints
   - Add business logic in services
   - Update repository interfaces if needed

3. **Testing**:
   - Write unit tests for services
   - Integration tests for handlers
   - Database tests for repositories

## 🚀 Deployment

### Production Deployment

1. **Environment Setup**
   ```bash
   # Create production .env
   cp .env.example .env.prod
   # Update with production values
   ```

2. **Build and Deploy**
   ```bash
   # Build production image
   docker-compose -f compose.yml build

   # Deploy with production settings
   docker-compose -f compose.yml up -d
   ```

3. **Health Monitoring**
   ```bash
   # Check application health
   curl https://your-domain.com/ping

   # Check container status
   docker-compose ps
   ```

### Docker Configuration

The project includes optimized Docker configurations:

- **Multi-stage Build**: Minimal production image
- **Health Checks**: Built-in health monitoring
- **Security**: Non-root user, minimal attack surface
- **Development Tools**: Air, dlv, migrate, sqlc included in dev image

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Application environment | `development` |
| `APP_PORT` | Server port | `8080` |
| `POSTGRES_URL` | PostgreSQL connection string | `postgres://user:password@db:5432/mydatabase?sslmode=disable` |
| `REDIS_URL` | Redis connection string | `redis://redis:6379/0` |
| `SECRET_KEY` | JWT secret key | `yourverysecretkey` |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go conventions and best practices
- Write meaningful commit messages
- Add tests for new features
- Update documentation as needed
- Ensure all tests pass before submitting PR

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [SQLC](https://sqlc.dev/) for excellent SQL code generation
- [Gin](https://gin-gonic.com/) for the high-performance web framework
- [PostgreSQL](https://www.postgresql.org/) for robust database functionality
- [Docker](https://www.docker.com/) for containerization support

