package model

import (
	"ecommerce-backend/internal/constant"
	"time"
)

type Product struct {
	ID          uint                    `json:"id" gorm:"primaryKey"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Stock       int                     `json:"stock"`
	ImageURL    string                  `json:"image_url"`
	Status      constant.ProductsStatus `json:"status"`
	CreatedAt   time.Time               `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time               `gorm:"autoUpdateTime" json:"updated_at"`
}
