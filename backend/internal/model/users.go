package model

import (
	"ecommerce-backend/internal/constant"
	"time"
)

type User struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	Role      constant.UserRole `json:"role"`
	CreatedAt time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
}
