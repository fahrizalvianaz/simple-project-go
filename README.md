# User Authentication and Management Service with JWT Support

A robust Go-based user authentication and management service that provides secure user registration, login, and profile management with JWT-based authentication. The service offers a RESTful API interface with standardized response handling and comprehensive test coverage.

This service is built using modern Go practices and industry-standard libraries including Gin for HTTP routing, GORM for database operations, and JWT for secure authentication. It features a clean architecture with clear separation of concerns between handlers, services, and repositories, making it both maintainable and extensible.

Key features include:
- Secure user registration with password hashing
- JWT-based authentication with configurable token expiration
- Protected profile endpoints
- PostgreSQL database integration with automatic migrations
- Comprehensive test coverage with mocking
- Standardized API response formatting
- Environment-based configuration management

## Repository Structure
```
.
├── cmd/
│   └── main.go                 # Application entry point
├── configs/
│   └── config.go               # Configuration management
├── internal/
│   └── users/                  # User domain implementation
│       ├── api/                # HTTP handlers and DTOs
│       ├── user.model.go       # User domain model
│       ├── user.repository.go  # Database operations
│       └── user.service.go     # Business logic
├── middleware/
│   └── middleware.go           # JWT authentication middleware
├── migrations/
│   └── migrations.go           # Database migration handler
├── pkg/                        # Shared utilities
│   ├── config.db.go           # Database connection
│   ├── generateToken.go       # JWT token generation
│   └── genericResponse.go     # API response formatting
└── test/                      # Test suites
    ├── handler/               # Handler tests
    ├── repository/           # Repository tests
    └── service/              # Service tests
```

## Usage Instructions
### Prerequisites
- Go 1.16 or higher
- PostgreSQL 12 or higher
- Environment variables configured in `.env` file

### Installation
1. Clone the repository:
```bash
git clone <repository-url>
cd <repository-name>
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables in `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
SECRET_KEY=your_secret_key
TOKEN_ISSUER=your_issuer
TOKEN_AUDIENCE=your_audience
```

### Quick Start
1. Start the server:
```bash
go run cmd/main.go
```

2. Register a new user:
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123","email":"test@example.com","name":"Test User"}'
```

3. Login to get JWT token:
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

### More Detailed Examples
1. Get user profile (protected endpoint):
```bash
curl http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer <your_jwt_token>"
```

### Troubleshooting
1. Database Connection Issues
- Error: "Failed to connect to database"
- Solution: 
  1. Verify PostgreSQL is running
  2. Check database credentials in `.env`
  3. Ensure database exists and is accessible

2. JWT Authentication Issues
- Error: "invalid or expired token"
- Solution:
  1. Verify token hasn't expired (default 24h)
  2. Ensure correct token format: `Bearer <token>`
  3. Check SECRET_KEY matches the one used for token generation

## Data Flow
The service follows a three-tier architecture with clear separation between API handlers, business logic, and data access.

```ascii
Client Request → JWT Middleware → Handler → Service → Repository → Database
     ↑                                                               ↓
     └───────────────── Response ← Handler ← Service ← Repository ←──┘
```

Component interactions:
1. JWT Middleware validates authentication token and extracts user context
2. Handlers validate request data and convert to domain models
3. Service layer implements business logic and orchestrates operations
4. Repository layer handles database operations using GORM
5. Standardized response formatting for consistent API responses
6. Error handling at each layer with appropriate HTTP status codes
7. Database transactions managed at repository level

## Infrastructure

![Infrastructure diagram](./docs/infra.svg)
### Database Resources
- Table: `users`
  - Primary key: `id`
  - Unique constraints: `username`, `email`
  - Soft delete support via `deleted_at`

### Migrations
The service automatically runs migrations on startup to ensure database schema is up to date:
```go
migrations.Migrate(db)
```