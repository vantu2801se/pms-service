# 🛒 Product Manager System
Interview Assignment: Build an Online Store Product Management System

## Overview
- PMS is a simple product management system providing 4 main APIs:
    - GET Category: Retrieve all product categories.
    - POST Product: Create a new product.
    - GET Products: Retrieve the list of products.
    - GET Product Count by Category: Get the number of products grouped by category.

- The application is lightweight, easy to deploy, and ready for future extensions.

## Technologies and Tools
- Golang: Main programming language.
- Gin: HTTP web framework for Go.
- GORM (v2): ORM library for Go.
- MySQL: Database management system.
- Docker: Application containerization.
- Docker Compose: Managing multi-container Docker applications.

## Setup Instructions
### Requirements
- Docker
- Docker Compose

### Steps to Setup
- Clone the project

```bash
git clone https://github.com/vantu2801se/pms-service
cd product-manager-system
```
- Run service
```bash
docker-compose up --build
```
- Service listening to port: 
    - PMS API Server: localhost:8000
    - MySQL Database: localhost:3306

Note: Docker Compose will automatically initialize the database and apply migrations if configured.

## API Documentation
Note: Detailed API documentation (requests & responses) will be updated later.

Available APIs:

GET /categories

POST /products

GET /products?offset={offset}&limit={limit}

GET /products/count

## Testing APIs
You can use Postman or cURL to test APIs.
- Example for GET /v1/categories:
```bash
curl -X GET http://localhost:8000/v1/categories
```

- Example for POST /v1/product
```bash
curl -X POST http://localhost:8000/v1/product \
-H "Content-Type: application/json" \
-d '{
    "product_name": "product_1",
    "description": "description_1",
    "price": 10.5,
    "quantity": 100,
    "category_id": 1
}'
```

- Example for GET /v1/products

```bash
curl -X GET http://localhost:8000/v1/products?offset=0&limit=10

curl -X GET http://localhost:8000/v1/products?status=in_stock&offset=0&limit=10

curl -X GET http://localhost:8000/v1/products?status=in_stock&status=out_of_stock&offset=0&limit=10
```

- Example for GET /v1/products/count

```bash
curl -X GET http://localhost:8000/v1/products/count
```

#### Reference: https://documenter.getpostman.com/view/16412669/2sB2j1gBhe
