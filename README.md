# Bookstore API Framework

A robust, Go-based RESTful API framework for managing bookstore operations with authentication, user management, and comprehensive CRUD operations.

## 📋 Table of Contents

- Features
- Project Structure
- Installation
- Configuration
- API Documentation
- Authentication
- Database
- Deployment
- Testing
- Contributing

## ✨ Features

- **User Management:** Registration, authentication, and profile management
- **JWT Authentication:** Secure route protection with JSON Web Tokens
- **RESTful API Design:** Clean API architecture following REST principles
- **Database Integration:** GORM-based PostgreSQL integration with migrations
- **Error Handling:** Consistent error responses and logging
- **Environment Configuration:** Environment-based configuration management
- **Middleware Support:** Request logging, authentication, CORS, etc.

## 🏗️ Project Structure

```
bookstore-framework/
├── .env                      # Environment variables (not tracked in git)
├── .env.example              # Example environment configuration
├── configs/                  # Configuration setup
│   └── config.go             # Application configuration
├── internal/                 # Internal application packages
│   ├── users/                # User domain
│   │   ├── api/              # API layer
│   │   │   ├── dto/          # Data Transfer Objects
│   │   │   ├── user.handler.go # HTTP handlers
│   │   │   └── user.router.go  # Route definitions
│   │   ├── user.model.go     # User data model
│   │   ├── user.repository.go # Data access layer
│   │   └── user.service.go   # Business logic layer
│   └── books/                # Books domain (similar structure)
├── middleware/               # HTTP middlewares
│   └── middleware.go         # JWT authentication and other middlewares
├── migrations/               # Database migrations
│   └── migrations.go         # Auto-migration setup
├── pkg/                      # Shared utility packages
│   ├── config.db.go          # Database connection setup
│   ├── generateToken.go      # JWT token generation
│   └── genericResponse.go    # Standardized response format
├── main.go                   # Application entry point
└── README.md                 # Project documentation
```

## 🚀 Installation

### Prerequisites

- Go 1.18 or later
- PostgreSQL 12 or later
- Git

### Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/bookstore-framework.git
   cd bookstore-framework
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## ⚙️ Configuration

### Environment Variables

```
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=bookstore

# JWT
SECRET_KEY=your-secret-key
TOKEN_ISSUER=bookstore-framework-api
TOKEN_AUDIENCE=bookstore-clients
```

## 📚 API Documentation

### Authentication Endpoints

#### Register a new user
```
POST /api/users/register
```
**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword",
  "full_name": "John Doe"
}
```

#### User Login
```
POST /api/users/login
```
**Request Body:**
```json
{
  "username": "johndoe",
  "password": "securepassword"
}
```
**Response:**
```json
{
  "code": 200,
  "message": "Login successful",
  "status": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "full_name": "John Doe",
      "created_at": "2025-03-06T13:45:22Z"
    }
  }
}
```

#### Get User Profile (Protected Route)
```
GET /api/users/profile
```
**Headers:**
```
Authorization: Bearer your-jwt-token
```

## 🔐 Authentication

The API uses JSON Web Tokens (JWT) for authentication:

1. **How it works:**
   - User logs in with credentials and receives a JWT token
   - Token must be included in the `Authorization` header as `Bearer token` for protected routes
   - Token contains encoded user information and an expiration time

2. **JWT Claims:**
   - `user_id`: Unique user identifier
   - `username`: User's username
   - `email`: User's email
   - Standard claims: `exp` (expiration), `iat` (issued at), `nbf` (not before)

3. **Protected Routes:**
   - All routes requiring authentication are protected by the `middleware.JWTAuth()` middleware
   - Middleware validates token signature, expiration, and extracts user information

## 💾 Database

### Migrations

The application uses GORM's AutoMigrate feature to manage database schema:

```go
// Running migrations
func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &users.User{},
        // Other models
    )
}
```
