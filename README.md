# Example-CRUD-Vue-3-FastAPI
- The example shows how to Building a Vue CRUD App with a FastAPI and using MySQL as a database.
- The article of this repository https://blog.stackpuz.com/building-a-vue-crud-app-with-a-fastapi
- To find more resources, please visit https://stackpuz.com

## Prerequisites
- Node.js
- Python 3.10
- MySQL

## Installation
- Clone this repository `git clone https://github.com/stackpuz/Example-CRUD-Vue-3-FastAPI .`
- Change directory to Vue project. `cd view`
- Install the Vue dependencies. `npm install`
- Change directory to FastAPI project. `cd ../api`
- Activate virtual environment and install packages. `pip install -r requirements.txt`
- Create a new database and run [/database.sql](/database.sql) script to create tables and import data.
- Edit the database configuration in [/api/.env](/api/.env) file.

## Run project

- Run Vue project. `npm run dev`
- Run FastAPI project `uvicorn app.main:app`
- Navigate to http://localhost:5173

## Run with Docker Compose

You can run the entire stack (backend(s), frontend, database, tracing, and logging) using Docker Compose. This will start:
- MySQL database
- Go API backend
- FastAPI backend
- Vue 3 frontend
- Jaeger (distributed tracing UI)
- VictoriaLogs (log storage)
- Filebeat (log shipper)

1. **Start all services:**
   ```bash
   docker-compose -f docker-compose-go.yaml up --build
   ```
   This will build and start all containers. The first run may take a few minutes.

2. **Stop all services:**
   ```bash
   docker-compose -f docker-compose-go.yaml down
   ```

### Exposed Ports and UIs

| Service         | URL/Port                | Description                       |
|----------------|-------------------------|-----------------------------------|
| Frontend (UI)  | http://localhost:5173   | Product CRUD web interface        |
| Go API         | http://localhost:8000   | REST API (for frontend/clients)   |
| FastAPI        | http://localhost:8001   | REST API (alternative backend)    |
| MySQL          | localhost:3306          | MySQL database                    |
| Jaeger UI      | http://localhost:16686  | Distributed tracing dashboard     |
| VictoriaLogs   | http://localhost:9428   | Log storage/query UI              |

- **Product CRUD UI:**
  - Access the web interface at [http://localhost:5173](http://localhost:5173) to manage products.
- **Jaeger UI:**
  - View distributed traces at [http://localhost:16686](http://localhost:16686)
- **VictoriaLogs UI:**
  - Query and explore logs at [http://localhost:9428](http://localhost:9428)