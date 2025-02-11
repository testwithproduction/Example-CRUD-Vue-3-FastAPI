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