package model

import "time"

type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `json:"cart_id"`
	ProductID uint      `json:"product_id"` // สินค้าที่ถูกเพิ่มใน cart
	Quantity  int       `json:"quantity"`   // จำนวนสินค้า
	CreatedAt time.Time `json:"created_at"`
}
