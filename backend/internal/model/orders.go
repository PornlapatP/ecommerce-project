package model

import (
	"ecommerce-backend/internal/constant"
	"time"
)

type Order struct {
	ID          uint                 `gorm:"primaryKey"`
	UserID      uint                 `json:"user_id"`
	TotalAmount float64              `json:"total_amount"`
	Status      constant.OrderStatus `json:"status" gorm:"default:'pending'"`
	CreatedAt   time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
	Items       []OrderItem          `gorm:"foreignKey:OrderID"`
}
