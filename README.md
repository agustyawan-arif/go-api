# Go Products API

This is a simple RESTful API built with [Gin](https://github.com/gin-gonic/gin) in Go, providing basic CRUD operations for managing a list of products. The API allows clients to **Create**, **Read**, **Update**, and **Delete** (CRUD) products.

## Features

- **GET /products**: Retrieve all products.
- **GET /products/:id**: Retrieve a specific product by its ID.
- **POST /products**: Add a new product to the list.
- **PUT /products/:id**: Update an existing product by its ID.
- **DELETE /products/:id**: Delete a product by its ID.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.
- A terminal or command-line interface.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/agustyawan-arif/go-products-api.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-products-api
   ```
3. Install the dependencies:
   ```bash
   go mod tidy
   ```

## Running the API

1. Start the API server by running the following command in the project directory:
   ```bash
   go run main.go
   ```
2. The API will be available at `http://localhost:8080`.

## API Endpoints

### Get All Products

- **URL**: `/products`
- **Method**: `GET`
- **Description**: Retrieves a list of all available products.

```bash
curl -X GET http://localhost:8080/products
```

### Get Product by ID

- **URL**: `/products/:id`
- **Method**: `GET`
- **Description**: Retrieves a product by its ID.

```bash
curl -X GET http://localhost:8080/products/1
```

### Create a New Product

- **URL**: `/products`
- **Method**: `POST`
- **Description**: Adds a new product.
- **Request Body** (JSON):

```json
{
  "id": "3",
  "name": "Tablet",
  "price": 300.00,
  "quantity": 20
}
```

Example:

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"id": "3", "name": "Tablet", "price": 300.00, "quantity": 20}'
```

### Update a Product

- **URL**: `/products/:id`
- **Method**: `PUT`
- **Description**: Updates an existing product.
- **Request Body** (JSON):

```json
{
  "id": "3",
  "name": "Tablet Pro",
  "price": 400.00,
  "quantity": 30
}
```

Example:

```bash
curl -X PUT http://localhost:8080/products/3 \
-H "Content-Type: application/json" \
-d '{"id": "3", "name": "Tablet Pro", "price": 400.00, "quantity": 30}'
```

### Delete a Product

- **URL**: `/products/:id`
- **Method**: `DELETE`
- **Description**: Deletes a product by its ID.

```bash
curl -X DELETE http://localhost:8080/products/3
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
