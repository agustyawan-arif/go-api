package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Product represents a product model
type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products = []Product{
	{ID: "1", Name: "Laptop", Price: 1000.50, Quantity: 10},
	{ID: "2", Name: "Smartphone", Price: 500.99, Quantity: 50},
}

// getProducts handles the GET /products endpoint
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// getProductByID handles the GET /products/:id endpoint
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.IndentedJSON(http.StatusOK, product)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// createProduct handles the POST /products endpoint
func createProduct(c *gin.Context) {
	var newProduct Product

	// Bind JSON to the newProduct object
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new product to the products slice
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

// updateProduct handles the PUT /products/:id endpoint
func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.IndentedJSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// deleteProduct handles the DELETE /products/:id endpoint
func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			// Delete product from slice
			products = append(products[:i], products[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func main() {
	router := gin.Default()

	// API routes
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	// Start the server on port 8080
	router.Run("localhost:8080")
}
