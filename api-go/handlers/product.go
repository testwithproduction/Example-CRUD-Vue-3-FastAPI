package handlers

import (
	"net/http"

	"api-go/config"
	"api-go/models"

	"log/slog"

	"github.com/gin-gonic/gin"
)

// GetProducts returns all products
func GetProducts(c *gin.Context) {
	slog.Debug("GetProducts handler called")
	var products []models.Product
	result := config.DB.WithContext(c.Request.Context()).Find(&products)
	if result.Error != nil {
		slog.Error("Failed to fetch products", "err", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	slog.Debug("Fetched products", "count", len(products))
	c.JSON(http.StatusOK, products)
}

// GetProduct returns a single product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	slog.Debug("GetProduct handler called", "id", id)
	var product models.Product
	result := config.DB.WithContext(c.Request.Context()).First(&product, id)
	if result.Error != nil {
		slog.Error("Product not found", "id", id, "err", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	slog.Debug("Fetched product", "id", id)
	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product
func CreateProduct(c *gin.Context) {
	slog.Debug("CreateProduct handler called")
	var productCreate models.ProductCreate
	if err := c.ShouldBindJSON(&productCreate); err != nil {
		slog.Error("Failed to bind JSON for product creation", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Debug("Creating product", "name", productCreate.Name, "price", productCreate.Price)
	product := models.Product{
		Name:  productCreate.Name,
		Price: productCreate.Price,
	}
	result := config.DB.WithContext(c.Request.Context()).Create(&product)
	if result.Error != nil {
		slog.Error("Failed to create product", "err", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	slog.Debug("Product created", "id", product.ID)
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates an existing product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	slog.Debug("UpdateProduct handler called", "id", id)
	var product models.Product
	// Check if product exists
	result := config.DB.WithContext(c.Request.Context()).First(&product, id)
	if result.Error != nil {
		slog.Error("Product not found for update", "id", id, "err", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	var productUpdate models.ProductUpdate
	if err := c.ShouldBindJSON(&productUpdate); err != nil {
		slog.Error("Failed to bind JSON for product update", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Debug("Updating product", "id", id, "name", productUpdate.Name, "price", productUpdate.Price)
	// Update product
	product.Name = productUpdate.Name
	product.Price = productUpdate.Price
	result = config.DB.WithContext(c.Request.Context()).Save(&product)
	if result.Error != nil {
		slog.Error("Failed to update product", "id", id, "err", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	slog.Debug("Product updated", "id", id)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	slog.Debug("DeleteProduct handler called", "id", id)
	result := config.DB.WithContext(c.Request.Context()).Delete(&models.Product{}, id)
	if result.Error != nil {
		slog.Error("Failed to delete product", "id", id, "err", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		slog.Warn("Product not found for deletion", "id", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	slog.Debug("Product deleted", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
} 