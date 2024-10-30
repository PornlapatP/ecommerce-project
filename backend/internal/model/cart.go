package model

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"user_id"`  // User ID ที่เป็นเจ้าของ cart นี้
	Products  []CartItem `json:"products"` // รายการสินค้าใน cart
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
