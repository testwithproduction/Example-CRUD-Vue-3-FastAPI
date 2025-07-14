package models

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:50;not null"`
	Price     float64        `json:"price" gorm:"type:decimal(12,2);not null"`
}

// TableName overrides the default table name for Product
func (Product) TableName() string {
	return "Product"
}

type ProductCreate struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type ProductUpdate struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
} 