package services

import (
	"errors"
	"strconv"

	"mcp-server/database"
	"mcp-server/models"
)

// ProductService handles all product-related operations
type ProductService struct{}

// NewProductService creates a new product service instance
func NewProductService() *ProductService {
	return &ProductService{}
}

// GetAllProducts retrieves all products from the database
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// GetProductByID retrieves a product by its ID
func (s *ProductService) GetProductByID(idStr string) (*models.Product, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return nil, errors.New("product not found")
	}

	return &product, nil
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(name string, price float64) (*models.Product, error) {
	if name == "" {
		return nil, errors.New("product name is required")
	}

	if price <= 0 {
		return nil, errors.New("product price must be greater than 0")
	}

	product := models.Product{
		Name:  name,
		Price: price,
	}

	result := database.DB.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(idStr string, name string, price float64) (*models.Product, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return nil, errors.New("product not found")
	}

	if name == "" {
		return nil, errors.New("product name is required")
	}

	if price <= 0 {
		return nil, errors.New("product price must be greater than 0")
	}

	product.Name = name
	product.Price = price

	result = database.DB.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// DeleteProduct deletes a product by its ID
func (s *ProductService) DeleteProduct(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return errors.New("invalid product ID")
	}

	result := database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
} 