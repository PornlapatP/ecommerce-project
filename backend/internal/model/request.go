package model

import (
	"ecommerce-backend/internal/constant"
)

type RequestLogin struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
	// Surname  string `binding:"required"`
	// Email string `binding:"required"`
}

type RequestCreateUser struct {
	Username string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type RequestGetUserByID struct {
	ID uint `json:"id"`
}

type RequestUpdateRole struct {
	Role constant.UserRole
}

type RequestCreateProduct struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	ImageURL    string  `json:"imageurl"`
}

type RequestGetProduct struct {
	Status []constant.ProductsStatus `form:"status" validate:"dive,oneof=active inactive"`
}

type RequestGetProductById struct {
	ID uint `json:"id"`
}
type RequestUpdateProduct struct {
	Status constant.ProductsStatus
}

type RequestCreateOrder struct {
	UserID      uint                  `json:"user_id"`
	Items       []OrderProductRequest `json:"items"`
	TotalAmount float64               `json:"total_amount"`
}

type OrderProductRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type RequestGetOrder struct {
	Status []constant.OrderStatus `form:"status" validate:"dive,oneof=pending approved"`
}

type RequestGetOrderById struct {
	ID uint `json:"id"`
}

//	type RequestUpdateOrder struct {
//		Status constant.OrderStatus
//	}
type RequestUpdateOrder struct {
	TotalAmount float64 `json:"total_amount"` // ยอดรวมคำสั่งซื้อ
	Status      constant.OrderStatus
	Items       []OrderItemRequest `json:"items"` // รายการสินค้าในคำสั่งซื้อ (ถ้าต้องการอัปเดตรายการสินค้า)
}
type OrderItemRequest struct {
	ProductID uint `json:"product_id"` // ID ของสินค้า
	Quantity  int  `json:"quantity"`   // จำนวนสินค้า
}

// OrderItemRequest ใช้สำหรับรับข้อมูลจาก client
// type OrderItemRequest struct {
// 	ProductID uint `json:"product_id"`
// 	Quantity  int  `json:"quantity"`
// }

// OrderItem ใช้สำหรับการจัดการข้อมูลคำสั่งซื้อที่แท้จริง
type RequestUpdateStatusOrder struct {
	Status constant.OrderStatus
}

type RequestCreateCart struct {
	UserID   uint             `json:"user_id"`
	Products []RequestProduct `json:"products"`
}
type RequestProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
