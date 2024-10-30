package model

import "time"

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Product   Product   `gorm:"foreignKey:ProductID"` // เชื่อมโยงกับ Product
}
