# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a full-stack CRUD application demonstrating Vue 3 + FastAPI integration with MySQL database. The project consists of two main components:

- **Frontend**: Vue 3 application with Vite build system (in `/view` directory)
- **Backend**: FastAPI REST API with SQLAlchemy ORM (in `/api` directory)

## Development Commands

### Frontend (Vue 3)
```bash
cd view
npm install          # Install dependencies
npm run dev          # Start development server (http://localhost:5173)
npm run build        # Build for production
npm run preview      # Preview production build
```

### Backend (FastAPI)
```bash
cd api
pip install -r requirements.txt    # Install Python dependencies
uvicorn app.main:app               # Start API server (default: http://localhost:8000)
uvicorn app.main:app --reload      # Start with auto-reload for development
```

### Database Setup
1. Create MySQL database
2. Run `/database.sql` script to create tables and import data
3. Configure database connection in `/api/.env` file

## Architecture

### Backend Structure (`/api`)
- **`app/main.py`**: FastAPI application entry point with CORS middleware
- **`app/db.py`**: Database configuration and SQLAlchemy setup
- **`app/models/`**: SQLAlchemy ORM models (Product model)
- **`app/schemas/`**: Pydantic schemas for request/response validation
- **`app/routers/`**: API route handlers (RESTful endpoints for products)

### Frontend Structure (`/view`)
- **`src/App.vue`**: Main application component
- **`src/router.js`**: Vue Router configuration
- **`src/http.js`**: Axios HTTP client configuration
- **`src/components/product/`**: Product CRUD components (Index, Create, Edit, Delete, Detail)
- **`src/components/product/Service.js`**: API service layer for product operations

### API Endpoints
The FastAPI backend exposes RESTful endpoints under `/api` prefix:
- `GET /api/products` - List products
- `GET /api/products/{id}` - Get product by ID
- `POST /api/products` - Create new product
- `PUT /api/products/{id}` - Update product
- `DELETE /api/products/{id}` - Delete product

### Data Flow
1. Vue components use Service.js to make HTTP requests
2. Service.js uses http.js (Axios) to communicate with FastAPI
3. FastAPI routes handle requests and interact with MySQL via SQLAlchemy
4. Responses flow back through the same chain

## Docker Development

### Docker Compose Commands
```bash
docker-compose up --build    # Build and start all services
docker-compose up -d         # Start services in background
docker-compose down          # Stop and remove containers
docker-compose logs api      # View API logs
docker-compose logs frontend # View frontend logs
```

### Services
- **mysql**: MySQL 8.0 database (port 3306)
- **api**: FastAPI backend (port 8000)
- **frontend**: Vue 3 frontend (port 5173)

### Docker Environment
- Database automatically initializes with `/database.sql`
- API uses Python 3.13
- Frontend uses Node.js 20
- All services have health checks and proper dependencies

## Environment Configuration

- Frontend runs on port 5173 (Vite default)
- Backend runs on port 8000 (FastAPI default)
- Database configuration is stored in `/api/.env`
- CORS is configured to allow all origins for development