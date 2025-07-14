# MCP Product Server

A Model Context Protocol (MCP) server written in Go using [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) that provides CRUD operations for products, similar to the existing API but designed for MCP tool calls.

## Features

- **Product Management**: Full CRUD operations for products
- **SQLite Database**: Lightweight database storage
- **MCP Protocol**: Standards-compliant MCP implementation using mark3labs/mcp-go
- **JSON Responses**: Structured data responses
- **Streamable HTTP**: Modern HTTP transport with SSE support
- **Tool Schema Validation**: Automatic argument validation and type checking

## Available Tools

The MCP server exposes the following tools:

### 1. `list_products`
Retrieves all products from the database.

**Arguments**: None

**Example Response**:
```json
{
  "content": [
    {
      "type": "text",
      "text": "Found 2 products:\n```json\n[\n  {\n    \"id\": 1,\n    \"name\": \"Laptop\",\n    \"price\": 999.99,\n    \"created_at\": \"2024-01-01T00:00:00Z\",\n    \"updated_at\": \"2024-01-01T00:00:00Z\"\n  }\n]\n```"
    }
  ]
}
```

### 2. `get_product`
Retrieves a specific product by ID.

**Arguments**:
- `id` (string): Product ID

**Example Request**:
```json
{
  "name": "get_product",
  "arguments": {
    "id": "1"
  }
}
```

### 3. `create_product`
Creates a new product.

**Arguments**:
- `name` (string): Product name (required)
- `price` (string): Product price as string (required, must be > 0)

**Example Request**:
```json
{
  "name": "create_product",
  "arguments": {
    "name": "Smartphone",
    "price": "599.99"
  }
}
```

### 4. `update_product`
Updates an existing product.

**Arguments**:
- `id` (string): Product ID (required)
- `name` (string): New product name (required)
- `price` (string): New product price as string (required, must be > 0)

**Example Request**:
```json
{
  "name": "update_product",
  "arguments": {
    "id": "1",
    "name": "Updated Laptop",
    "price": "1099.99"
  }
}
```

### 5. `delete_product`
Deletes a product by ID.

**Arguments**:
- `id` (string): Product ID

**Example Request**:
```json
{
  "name": "delete_product",
  "arguments": {
    "id": "1"
  }
}
```

## Installation

1. **Clone the repository**:
   ```bash
   cd mcp-server
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Start MySQL database** (if not already running):
   ```bash
   # From the root directory
   docker-compose up mysql -d
   ```

4. **Run the server**:
   ```bash
   go run main.go
   ```

### Using Docker

1. **Start MySQL database** (if not already running):
   ```bash
   # From the root directory
   docker-compose up mysql -d
   ```

2. **Start the MCP server**:
   ```bash
   docker-compose -f docker-compose-mcp.yaml up --build
   ```

## Configuration

### Environment Variables

- `MCP_PORT`: Server port (default: 8080)
- `DB_USER`: MySQL username (default: crud_user)
- `DB_PASSWORD`: MySQL password (default: crud_password)
- `DB_HOST`: MySQL host (default: localhost)
- `DB_PORT`: MySQL port (default: 3306)
- `DB_DATABASE`: MySQL database name (default: crud_db)
- `DATABASE_URL`: Alternative full MySQL connection string

### Database

The server connects to the same MySQL database as your existing API. Make sure the MySQL service is running before starting the MCP server.

**Local Development:**
Create a `.env` file in the mcp-server directory:
```bash
DB_USER=crud_user
DB_PASSWORD=crud_password
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=crud_db
MCP_PORT=8080
```

**Docker:**
The MCP server will automatically connect to the MySQL service defined in your main docker-compose.yml.

## API Endpoints

### Tool Call Endpoint
- **URL**: `POST /mcp`
- **Content-Type**: `application/json`
- **Body**: JSON-RPC 2.0 compliant tool call request

### Example Request Format
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "create_product",
    "arguments": {
      "name": "Laptop",
      "price": 999.99
    }
  }
}
```

## Usage Examples

### Using curl

1. **List all products**:
   ```bash
   curl -X POST http://localhost:8080/mcp \
     -H "Content-Type: application/json" \
     -d '{
       "jsonrpc": "2.0",
       "id": 1,
       "method": "tools/call",
       "params": {
         "name": "list_products",
         "arguments": {}
       }
     }'
   ```

2. **Create a product**:
   ```bash
   curl -X POST http://localhost:8080/mcp \
     -H "Content-Type: application/json" \
     -d '{
       "jsonrpc": "2.0",
       "id": 1,
       "method": "tools/call",
       "params": {
         "name": "create_product",
         "arguments": {
           "name": "Tablet",
           "price": 299.99
         }
       }
     }'
   ```

3. **Get a product**:
   ```bash
   curl -X POST http://localhost:8080/mcp \
     -H "Content-Type: application/json" \
     -d '{
       "jsonrpc": "2.0",
       "id": 1,
       "method": "tools/call",
       "params": {
         "name": "get_product",
         "arguments": {
           "id": "1"
         }
       }
     }'
   ```

### Using with MCP Clients

The server is designed to work with MCP-compatible clients. Configure your MCP client to connect to:

```
http://localhost:8080/mcp
```

### Testing

Run the test client example:

```bash
go run test_client_example.go
```

## Project Structure

```
mcp-server/
├── main.go                    # Application entry point with mcp-go
├── go.mod                     # Go module file
├── database/
│   └── database.go            # Database configuration
├── models/
│   └── product.go             # Product data model
├── services/
│   └── product_service.go     # Business logic layer
├── test_client_example.go     # Example test client
└── README.md                  # This file
```

## Development

### Adding New Tools

1. Add the tool definition using `mcp.NewTool()` in `main.go`
2. Add the tool handler function using `s.AddTool()`
3. Use proper argument validation with `req.RequireString()`, `req.RequireFloat()`, etc.

### Database Schema Changes

1. Modify the model in `models/product.go`
2. The database will automatically migrate on startup

## Error Handling

The server provides detailed error messages for:
- Invalid tool names
- Missing required arguments
- Invalid data formats
- Database errors
- Product not found errors

## Logging

The server uses colored logging to distinguish between:
- 🚀 Server startup messages
- 🔧 Tool calls
- ✅ Successful operations
- ❌ Errors
- 📊 Arguments and responses

## License

This project is part of the Example-CRUD-Vue-3-FastAPI repository. 