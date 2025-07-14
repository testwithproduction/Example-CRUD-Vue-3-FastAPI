package handlers

import (
	"net/http"

	"api-go/config"
	"api-go/models"

	"github.com/gin-gonic/gin"
)

// GetProducts returns all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	result := config.DB.WithContext(c.Request.Context()).Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct returns a single product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	
	result := config.DB.WithContext(c.Request.Context()).First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product
func CreateProduct(c *gin.Context) {
	var productCreate models.ProductCreate
	
	if err := c.ShouldBindJSON(&productCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	product := models.Product{
		Name:  productCreate.Name,
		Price: productCreate.Price,
	}
	
	result := config.DB.WithContext(c.Request.Context()).Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates an existing product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	
	// Check if product exists
	result := config.DB.WithContext(c.Request.Context()).First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	var productUpdate models.ProductUpdate
	if err := c.ShouldBindJSON(&productUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Update product
	product.Name = productUpdate.Name
	product.Price = productUpdate.Price
	
	result = config.DB.WithContext(c.Request.Context()).Save(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	
	result := config.DB.WithContext(c.Request.Context()).Delete(&models.Product{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
} 