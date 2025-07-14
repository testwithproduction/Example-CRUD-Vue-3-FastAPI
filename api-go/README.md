# Go API - Product CRUD

This is a Go implementation of the Product CRUD API using Gin framework and GORM.

## Features

- **Product CRUD Operations**: Create, Read, Update, Delete products
- **MySQL Database**: Using GORM ORM
- **RESTful API**: Standard HTTP methods
- **CORS Support**: Cross-origin resource sharing enabled
- **Static File Serving**: Serves the frontend HTML file
- **Environment Configuration**: Configurable via .env file

## API Endpoints

```
GET    /api/products      - List all products
GET    /api/products/:id  - Get product by ID
POST   /api/products      - Create new product
PUT    /api/products/:id  - Update product
DELETE /api/products/:id  - Delete product
GET    /                 - Serve static index.html
```

## Prerequisites

- Go 1.21 or higher
- MySQL database
- Docker (optional)

## Environment Variables

Create a `.env` file with the following variables:

```env
DB_USER=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=crud_db
DATABASE_URL=  # Optional: full database URL
```

## Running the Application

### Local Development

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

3. **Build the application:**
   ```bash
   go build -o main .
   ./main
   ```

### Using Docker

1. **Build the Docker image:**
   ```bash
   docker build -t api-go .
   ```

2. **Run the container:**
   ```bash
   docker run -p 8080:8080 --env-file .env api-go
   ```

## Database Schema

The application automatically creates the Product table with the following structure:

```sql
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `price` decimal(12,2) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
);
```

## Project Structure

```
api-go/
├── main.go                 # Entry point
├── go.mod                  # Go module file
├── go.sum                  # Dependency checksums
├── .env                    # Environment variables
├── Dockerfile             # Container configuration
├── config/
│   └── database.go        # Database configuration
├── models/
│   └── product.go         # Product model
├── handlers/
│   └── product.go         # HTTP handlers
├── routes/
│   └── routes.go          # Route definitions
├── middleware/
│   └── cors.go            # CORS middleware
└── static/
    └── index.html         # Frontend file
```

## Dependencies

- **Gin**: HTTP web framework
- **GORM**: ORM library
- **MySQL Driver**: Database connectivity
- **CORS**: Cross-origin resource sharing
- **godotenv**: Environment variable loading

## Comparison with FastAPI

| Feature | FastAPI (Python) | Gin (Go) |
|---------|------------------|----------|
| Framework | FastAPI | Gin |
| ORM | SQLAlchemy | GORM |
| Validation | Pydantic | Struct tags |
| Async | async/await | Goroutines |
| Performance | Good | Excellent |
| Type Safety | Good | Excellent | 