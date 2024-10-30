package model

import (
	"ecommerce-backend/internal/constant"
	"time"
)

type BaseResponse[DataType any] struct {
	Message string   `json:"message,omitempty"`
	Data    DataType `json:"data,omitempty"`
}
type ResponseUser[DataType any] struct {
	Message   string    `json:"message,omitempty"`
	Data      DataType  `json:"data,omitempty"`
	Userid    uint      `json:"userid,omitempty"`    // ควรเป็น uint
	Username  string    `json:"username,omitempty"`  // ควรเป็น string
	Email     string    `json:"email,omitempty"`     // ควรเป็น string
	CreatedAt time.Time `json:"createdat,omitempty"` // ควรเป็น time.Time
}

type BaseResponseList[DataType any] struct {
	Count   int      `json:"count"`
	Results DataType `json:"results"`
}

type ResponseProduct struct {
	ID          uint                    `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Stock       int                     `json:"stock"`
	ImageURL    string                  `json:"image_url"`
	Status      constant.ProductsStatus `json:"status"`
	CreatedAt   time.Time               `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time               `gorm:"autoUpdateTime" json:"updated_at"`
}
type ResponseCart struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	Products  []CartItem `json:"products"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
